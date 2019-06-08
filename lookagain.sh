find . -name '*.sh' | sed  's#/##g' | sed  's/test//g' | cut -f2 -d '.' | cut -f2 -d 'h'
