package main

func Apply_dir_numpad(org rune,dir rune)rune{
    switch org{
    case 'A':
        switch dir{
        case '<':
            return '0'
        case '>':
            return ERROR
        case '^':
            return '3'
        case 'v':
            return ERROR
        }
    case '0':
        switch dir{
        case '<':
            return ERROR
        case '>':
            return 'A'
        case '^':
            return '2'
        case 'v':
            return ERROR
        }
    case '1':
        switch dir{
        case '<':
            return ERROR
        case '>':
            return '2'
        case '^':
            return '4'
        case 'v':
            return ERROR
        }
    case '2':
        switch dir{
        case '<':
            return '1'
        case '>':
            return '3'
        case '^':
            return '5'
        case 'v':
            return '0'
        }
    case '3':
        switch dir{
        case '<':
            return '2'
        case '>':
            return ERROR
        case '^':
            return '6'
        case 'v':
            return 'A'
        }
    case '4':
        switch dir{
        case '<':
            return ERROR
        case '>':
            return '5'
        case '^':
            return '7'
        case 'v':
            return '1'
        }
    case '5':
        switch dir{
        case '<':
            return '4'
        case '>':
            return '6'
        case '^':
            return '8'
        case 'v':
            return '2'
        }
    case '6':
        switch dir{
        case '<':
            return '5'
        case '>':
            return ERROR
        case '^':
            return '9'
        case 'v':
            return '3'
        }
    case '7':
        switch dir{
        case '<':
            return ERROR
        case '>':
            return '8'
        case '^':
            return ERROR
        case 'v':
            return '4'
        }
    case '8':
        switch dir{
        case '<':
            return '7'
        case '>':
            return '9'
        case '^':
            return ERROR
        case 'v':
            return '5'
        }
    case '9':
        switch dir{
        case '<':
            return '8'
        case '>':
            return ERROR
        case '^':
            return ERROR
        case 'v':
            return '6'
        }
    }
    return ERROR
}
func Apply_dir_dir(org rune,dir rune)rune{
    switch org{
    case 'A':
        switch dir{
        case '<':
            return '^'
        case '>':
            return ERROR
        case '^':
            return ERROR
        case 'v':
            return '>'
        }
    case '^':
        switch dir{
        case '<':
            return ERROR
        case '>':
            return 'A'
        case '^':
            return ERROR
        case 'v':
            return 'v'
        }
    case '<':
        switch dir{
        case '<':
            return ERROR
        case '>':
            return 'v'
        case '^':
            return ERROR
        case 'v':
            return ERROR
        }
    case 'v':
        switch dir{
        case '<':
            return '<'
        case '>':
            return '>'
        case '^':
            return '^'
        case 'v':
            return ERROR
        }
    case '>':
        switch dir{
        case '<':
            return 'v'
        case '>':
            return ERROR
        case '^':
            return 'A'
        case 'v':
            return ERROR
        }

    }
    return ERROR
}

func dir_to_dir(beg rune,end rune)seq{
    switch beg{
    case 'A':{
        switch end{
            case 'A':
                return seq{'A'}
            case '^':
                return seq{'<','A'}
            case '>':
                return seq{'v','A'}
            case 'v':
                return seq{'v','<','A'}
            case '<':
                return seq{'v','<','v','A'}
        }
    }
    case '^':{
        switch end{
            case 'A':
                return seq{'>','A'}
            case '^':
                return seq{'A'}
            case '>':
                return seq{'v','>','A'}
            case 'v':
                return seq{'v','A'}
            case '<':
                return seq{'v','<','A'}
        }
    }
    case '>':{
        switch end{
            case 'A':
                return seq{'^','A'}
            case '^':
                return seq{'<','^','A'}
            case '>':
                return seq{'A'}
            case 'v':
                return seq{'<','A'}
            case '<':
                return seq{'<','<','A'}
        }
    }
    case 'v':{
        switch end{
            case 'A':
                return seq{'>','^'}
            case '^':
                return seq{'^','A'}
            case '>':
                return seq{'>','A'}
            case 'v':
                return seq{'A'}
            case '<':
                return seq{'<','A'}
        }
    }
    case '<':{
        switch end{
            case 'A':
                return seq{'>','>','^','A'}
            case '^':
                return seq{'>','^','A'}
            case '>':
                return seq{'>','>','A'}
            case 'v':
                return seq{'>','A'}
            case '<':
                return seq{'A'}
        }
    }
}
return nil
}
