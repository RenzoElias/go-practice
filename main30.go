// pip - python
// go get -- golang

// go get golang.org/x/tour
// go get -v golang.org/x/tour // -v // nos da informacion de la instalacion
// go get -v -u golang.org/x/tour // -u // aunque este instalada en el gopath, igualmente lo vuelva a descargar

// go install -v golang.org/x/website/tour@latest

// Mensaje de error
// go: modules disabled by GO111MODULE=off; see 'go help modules'

// lo que hice fue introducir este comando:
// export GO111MODULE=on

// y luego ya pude descargarlo con el mismo comando😎:
// go install -v golang.org/x/website/tour@latest