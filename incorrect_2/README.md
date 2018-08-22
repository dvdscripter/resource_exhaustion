This code has a type of DoS (denial of service) called Resource Exhaustion.
Resource Exhaustion occurs when developers allocate some type of resource and 
forget to clean up after job is done.

http.Request.ParseMultipartForm(maxMemory int64) will create files at temporary 
directory if maxMemory is exceeded. After running this code with input csv you 
can check your temporary directory and one or more "multipart*" files will be 
there, even after closing app. Setting a limit in how much size temporary 
directory files can mitigate the problem for a while but another process can 
fill your limit quota and stopping to process user requests aren't a good 
solution.