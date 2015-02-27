package photoweb
import (
    "net/http"
    "io"
    "log"
    "os"
    "text/template"
    "io/ioutil"
    "path"
    "runtime/debug"
    "strings"
)

const (
    ListDir = 0x0001
    UPLOAD_DIR ="d:\\uploads"
    TEMPLATE_DIR ="photoweb/template"
)

var templates map[string]*template.Template = make(map[string]*template.Template)


func check(err error){
    if err!=nil{
        panic(err)
    }
}

func uploadHandler (w http.ResponseWriter, r *http.Request){
    if r.Method =="GET"{
        w.Header().Set("Content-Type","text/html")
        renderHtml(w,"upload",nil)
        return
    }
    if r.Method =="POST"{
        f,h,err:=r.FormFile("image")
        check(err)
        filename:=h.Filename
        defer f.Close()
        t,err:=os.Create(UPLOAD_DIR+"/"+filename)
        check(err)
        defer t.Close()
        _,err =io.Copy(t,f)
        check(err)
        http.Redirect(w,r,"/view?id="+filename,http.StatusFound)
    }
}

func viewHandler (w http.ResponseWriter, r *http.Request){
    id :=r.FormValue("id")
    imagePath :=UPLOAD_DIR+"/"+id
    if !isExists(imagePath){
        http.NotFound(w,r)
        return
    }

    w.Header().Set("Content-Type","image")
    http.ServeFile(w,r,imagePath)
}

func listHandler(w http.ResponseWriter, r *http.Request){
    fileInfoArr,err:=ioutil.ReadDir(UPLOAD_DIR)
    check(err)
    locals:=make(map[string]interface{})
    images :=[]string{}
    for _,fileInfo :=range fileInfoArr {
        images =append(images,fileInfo.Name())
    }
    locals["images"]=images
    renderHtml(w,"list",locals)
}

func renderHtml(w http.ResponseWriter,tmpl string,locals map[string]interface{}){
    err:=templates[tmpl].Execute(w,locals)
    check(err)
}

func isExists(path string) bool{
    _,err:=os.Stat(path)
    if err==nil{
        return true
    }
    return os.IsExist(err)
}

func safeHandler(fn http.HandlerFunc) http.HandlerFunc{
    return func(w http.ResponseWriter,r *http.Request){
        defer func(){
            if err,ok :=recover().(error);ok{
                http.Error(w,err.Error(),http.StatusInternalServerError)
                log.Println("Warn:panic in %v. - %v",fn,err)
                log.Println(string(debug.Stack()))
            }
        }()
        fn(w,r)
    }
}

func staticDirHandler(mux *http.ServeMux,prefix string,staticDir string,flags int){
    mux.HandleFunc(prefix,func(w http.ResponseWriter,r *http.Request){
        file :=staticDir+r.URL.Path[len(prefix)-1:]
        if flags & ListDir==0{
            if exist:=isExists(file);!exist{
                http.NotFound(w,r)
                return
            }
        }
        http.ServeFile(w,r,file)
    })
}

func initTmpl(){
    fileInfoArr,err :=ioutil.ReadDir(TEMPLATE_DIR)
    check(err)
    for _,fileInfo :=range fileInfoArr{
        tempName :=fileInfo.Name()
        if ext:=path.Ext(tempName);ext!=".html"{
            continue
        }
        tmpl:=strings.Split(tempName,".")[0]
        tempPath:=TEMPLATE_DIR+"/"+tempName
        log.Println("loading template:",tempPath)
        t :=template.Must(template.ParseFiles(tempPath))
        templates[tmpl]=t
    }
}

func Start(){
    initTmpl()
    mux:=http.NewServeMux()
    staticDirHandler(mux,"/public/","photoweb/public",0)
    mux.HandleFunc("/upload",safeHandler(uploadHandler))
    mux.HandleFunc("/",safeHandler(listHandler))
    mux.HandleFunc("/view",safeHandler(viewHandler))
    err :=http.ListenAndServe(":8080",mux)
    if err!=nil{
        log.Fatal("ListenAndServe:",err.Error())
    }
}