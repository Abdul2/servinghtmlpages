Go lang sample code showing how to serve dynamic and static content.
-----------------------------------------------------------------

Based on work of http://www.alexedwards.net/

`work in progress ..` 


## Folder structure


```

├── README.md
├── app.go
├── static
│   ├── example.html
│   └── stylesheets
│       └── main.css
└── templates
    ├── example.html
    └── layout.html
    
```

* GO app is in the root of the project 
* static folder servers static pages and stylesheets
* templates contains nested templates. the dynamic content is generated from the templates. 

## the logic 

* main function  

     * starts a server to listen on port 3000
     * enables the file system in the static folder to serve http calls
     * if the involves dynamic content, calls serveTemplates function   

* serveTemplates 

	* 
	
To do:
------

 - Include instructions on how to run the code and explain output
 - Tell the story and the logic of the code
 - link this sample to Wiki example and phonebook code samples to develop  a complete Go web app sample 


