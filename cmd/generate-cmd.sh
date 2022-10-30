# !/bin/bash

cmds=("add"
    "cat-file"    
    "checkout"    
    "commit"      
    "hash-object" 
    "log"         
    "ls-tree"     
    "merge"       
    "rebase"      
    "rev-parse"   
    "rm"               
    "show-ref"     
    "tag")

for cmd in ${cmds[@]}; do 
    # cat ./init.go | sed "s/init/$cmd/g" > $cmd.go
    echo "rootCmd.AddCommand($cmd""Cmd)" | sed "s/-//g"
done