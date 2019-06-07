



find . -name '*.sh' | sed  's#/##g' | cut -f2 -d '.' | cut -f2 -d 'h' | sed  's/"test"//g' 