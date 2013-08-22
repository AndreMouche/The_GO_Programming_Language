/**
 * Created with IntelliJ IDEA.
 * User: fun
 * Date: 13-8-22
 * Time: 下午2:22
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"io/ioutil"
	"html/template"
	"path"
	"runtime/debug"
)

const (
	UPLOAD_DIR = "/tmp/uploads"
	TEMPLATE_DIR = "/tmp/uploads/resources"
	ListDir = 0x0001
)

var templates = make(map[string]*template.Template)
func init() {
	fileInfoArr,err := ioutil.ReadDir(TEMPLATE_DIR)
	if err != nil {
		panic(err)
		return
	}

	var templateName,templatePath string
	for _,fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		if ext := path.Ext(templateName);ext != ".html" {
		    continue
		}
		templatePath = TEMPLATE_DIR + "/" + templateName
		log.Println("Loading template:",templatePath)
		t := template.Must(template.ParseFiles(templatePath))
		templates[templatePath] = t
	}

//
//	for _,temp := range []string{"upload","list"}  {
//		t := template.Must(template.ParseFiles(UPLOAD_DIR + "/resources/" + temp + ".html"))
//		templates[temp] = t
//	}
}


func uploadHandler(w http.ResponseWriter, r *http.Request) {
	 if r.Method == "GET" {
		 if err := readerHtml(w,TEMPLATE_DIR + "/upload.html",nil); err != nil {
			 http.Error(w,err.Error(),http.StatusInternalServerError)
			 return
		 }
//		 t,err := template.ParseFiles(UPLOAD_DIR + "/resources/upload.html")
//		 if err != nil {
//			 http.Error(w,err.Error(),http.StatusInternalServerError)
//			 return
//		 }
//		 t.Execute(w,nil)
//		 io.WriteString(w,  `<html xmlns="http://www.w3.org/1999/xhtml"><head></head><body>`+
//					         "<form method=\"POST\" action=\"/upload\" "+
//					         " enctype=\"multipart/form-data\">"+
//							 "Choose an image to upload: <input name=\"image\" type=\"file\" />"+
//							 "<input type=\"submit\" value=\"Upload\" />"+
//							 "</form></body>")
	     return
	 }

	 if r.Method == "POST" {
		 f,h,err := r.FormFile("image")
		 check(err)
//		 if err != nil {
//			 http.Error(w,err.Error(),http.StatusInternalServerError)
//			 return
//		 }

		 filename := h.Filename
		 defer f.Close()
		 t, err := os.Create(UPLOAD_DIR + "/" + filename)
		 check(err)
		  /*if err != nil {
			 http.Error(w,err.Error(),http.StatusInternalServerError)
			 return
		 }*/

		 defer t.Close()
		 if _, err := io.Copy(t,f);err != nil {
			 http.Error(w,err.Error(),http.StatusInternalServerError)
			 return
		 }
		 http.Redirect(w,r,"/view?id=" + filename,http.StatusFound)
	 }

}

func isExists(path string) bool {
	_,err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}
func viewHandler(w http.ResponseWriter, r *http.Request) {
	  imageId := r.FormValue("id")
	  imagePath := UPLOAD_DIR + "/" + imageId

	  if exists := isExists(imagePath);!exists {
		     http.NotFound(w,r)
		     return
      }
	  w.Header().Set("Content-Type","image")
	  http.ServeFile(w,r,imagePath)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir(UPLOAD_DIR)
	check(err)
	locals := make(map[string]interface{})
	images := []string{}
	for _,fileInfo := range fileInfoArr {
		images = append(images,fileInfo.Name())
	}
	locals["images"] = images
	log.Println("aa");
	if err := readerHtml(w,TEMPLATE_DIR + "/list.html",locals); err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
//	t,err := template.ParseFiles(UPLOAD_DIR + "/resources/list.html")
//	if err != nil {
//		http.Error(w,err.Error(),
//		http.StatusInternalServerError)
//		return
//	}
//
//	t.Execute(w,locals)

//
//	var listHtml string
//	for _,fileInfo := range fileInfoArr {
//		imageId := fileInfo.Name()
//		listHtml += "<li><a href=\"/view?id="+ imageId +"\">imgid</a></li>"
//	}
//	io.WriteString(w,"<ol>" + listHtml + "</ol>")
}



func readerHtml(w http.ResponseWriter,templ string,locals map[string]interface{}) error {

		return templates[templ].Execute(w,locals)

}

func safeHandler(fn http.HandlerFunc) http.HandlerFunc {      //引用闭包避免程序运行时出错崩溃
	 return func(w http.ResponseWriter, r *http.Request) {
		 defer func() {
			 if e, ok := recover().(error); ok {
				 http.Error(w,e.Error(),http.StatusInternalServerError)
				 log.Println("WARN:panic in %V -- %V",fn,e)
				 log.Println(string(debug.Stack()))
			 }
		 }()
		 fn(w,r)
	 }
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func staticDirHandler(mux *http.ServeMux,prefix string,staticDir string,flags int) {
	  mux.HandleFunc(prefix,func(w http.ResponseWriter,r *http.Request){
			  file := staticDir + r.URL.Path[len(prefix)-1:]
			  if (flags & ListDir) == 0{
				  if exists := isExists(file); !exists {
					  http.NotFound(w,r)
					  return
				  }
			  }
			  http.ServeFile(w,r,file)
		  })
}


func main(){
	mux := http.NewServeMux()
	staticDirHandler(mux,"/asserts/","./public",0)
	http.HandleFunc("/",safeHandler(listHandler))
	http.HandleFunc("/upload",safeHandler(uploadHandler))
	http.HandleFunc("/view",safeHandler(viewHandler) )

	err := http.ListenAndServe(":8089",nil)
	if err != nil {
		log.Fatal("ListenAndServe:",err.Error())
	}
}

