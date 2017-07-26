Grpc based golang implementation of the hangman game


TODO:
	```
	!!! show the user which incorrect letters have been already used

	maybe add config option in the menu  to specify total retry attempts limit etc
	remove duplicated "if OR" for the auto completer
	add and test the  grpc connection timeout and command timeout
	unit/integration tests
	api versioning in grpc ?
	vendor all dependancies
	```


	```

	// graficki dizajner je za ovo dobio goleme pare
		switch mistakes {
		case 0:
			art = `
	      __________
	     |/      |
	     |      
	     |     
	     |    
	     |   
	     |
	 ____|____`
		case 1:
			art = `
	      __________
	     |/      |
	     |      (_)
	     |     
	     |    
	     |   
	     |
	 ____|____`
		case 2:
			art = `
	      __________
	     |/      |
	     |      (_)
	     |       |
	     |       |
	     |       
	     |
	 ____|____`
		case 3:
			art = `
	      __________
	     |/      |
	     |      (_)
	     |      \|
	     |       |
	     |       
	     |
	 ____|____`
		case 4:
			art = `
	      __________
	     |/      |
	     |      (_)
	     |      \|/
	     |       |
	     |      
	     |
	 ____|____`
		case 5:
			art = `
	      __________
	     |/      |
	     |      (_)
	     |      \|/
	     |       |
	     |      /
	     |
	 ____|____`
		case 6:
			art = `
	      __________
	     |/      |
	     |      (_)
	     |      \|/
	     |       |
	     |      / \
	     |
	 ____|____`
		}
		return art

	```
