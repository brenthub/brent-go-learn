package photoweb
import (
    "net/http"
    "io"
    "log"
    "os"
    "text/template"
)

const (
    UPLOAD_DIR ="d:\\uploads"
)

func uploadHandler (w http.ResponseWriter, r *http.Request){
    if r.Method =="GET"{
        w.Header().Set("Content-Type","text/html")
        t,err:=template.ParseFiles("photoweb/template/upload.html")
        if err!=nil{
            http.Error(w,err.Error(),http.StatusInternalServerError)
            return
        }
        t.Execute(w,nil)
        return
    }
    if r.Method =="POST"{
        f,h,err:=r.FormFile("image")
        if err!=nil{
            http.Error(w,err.Error(),http.StatusInternalServerError)
            return
        }
        filename:=h.Filename
        defer f.Close()
        t,err:=os.Create(UPLOAD_DIR+"/"+filename)
        if(err!=nil){
            http.Error(w,err.Error(),http.StatusInternalServerError)
            return
        }
        defer t.Close()
        if _,err:=io.Copy(t,f); err!=nil{
            http.Error(w,err.Error(),http.StatusInternalServerError)
            return
        }
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

func Start(){
    http.HandleFunc("/upload",uploadHandler)
    http.HandleFunc("/view",viewHandler)
    err :=http.ListenAndServe(":8080",nil)
    if err!=nil{
        log.Fatal("ListenAndServe:",err.Error())
    }
}

func isExists(path string) bool{
    _,err:=os.Stat(path)
    if err==nil{
        return true
    }
    return os.IsExist(err)
}