send(req,req.url!,options)
DB.query(fmt.sprintf("SELECT * FROM books WHERE read = '%s'",r ))
