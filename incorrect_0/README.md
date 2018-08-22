This code has a type of DoS (denial of service) called Resource Exhaustion.
Resource Exhaustion occurs when developers allocate some type of resource and 
forget to clean up after job is done.

http.Request.ParseMultipartForm(maxMemory int64) will create files at temporary 
directory if maxMemory is exceeded. After running this code with input csv you 
can check your temporary directory is clean of almost any file. 
os.RemoveAll(os.TempDir()) seems like a cleaver solution at first glance but 
removing all files from temporary directory can break your own app or other 
process which also use temporary directory for processing. Also "multipart*" 
files can still be present because they aren't proper closed.