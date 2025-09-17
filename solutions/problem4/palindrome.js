function palindrome_product(){
    let n = 9911
    var ntoString = n.toString()
    let splitString = ntoString.split('')
    let revSplitString = splitString.reverse()
    let revString = revSplitString.join('')

    console.log(revString)
}

palindrome_product()