http.Request.ParseMultipartForm(maxMemory int64) will create files at temporary 
directory if maxMemory is exceeded. After running this code with input csv you 
can check your temporary directory is clean of any processing file. Go standard 
library will call http.Request.MultipartForm.RemoveAll() and clean any file 
created by ParseMultipartForm. You need to proper release the resource with 
Close() function.