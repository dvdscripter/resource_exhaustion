# Go service DoS by Resource Exhaustion

When working with multipart forms Go will try to store in memory the entire form 
and if size exceeds limit boundaries temporary files take place. `net/http` 
didn't warm developers `http.ListenAndServe` will try to clean up any created
temporary file and any working temporary file also need to be ready to remove.

If file is still in use (no Close() yet) `http.ListenAndServe` clean up can fail 
and your temporary directory can start pilling up. This is when resource 
exhaustion can happen by lack of inodes or disk space. Until Go decides to fix 
or not is developer responsibility to close any open files.

This repository has some examples of improper handling of problem and secure 
as well. Goal here is to explain the concept of not leaving any resource behind 
for OS/language to handle for you.

Secure Code Warrior has all rights about this challenge.
