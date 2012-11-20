goseg: A Chinese Word Segmentation Library in GoLang
==================================================

* goseg use max probability product to segment words, and use ANN to find new words

* goseg also provide a server which can handle POST request of sentence and return words segmented

* library usage


			sentence := "买水果然后去世博园"
			words := goseg.Cut([]rune(sentence))
    		for _,w := range words{
    			fmt.Println(w)
    		}

* goseg server usage

		go run goseg_server.go

		or 

		go build goseg_server.go
		./goseg_server


