console.log("Hello Euler!")

var outputElement = document.getElementById("Introduction")
outputElement.textContent = "How many times will we do this?"

let a = 5
var b = 6
const c = 12

let palindrome = 9009
let d = (c*a)/b
d.toString

palindrome.toString
var product = (a,b)=>a*b
console.log(product(3,5))

var myFamily = new Array("Senior", "Mother", "Mum Shiku", "Muthomi")
console.log(myFamily)
var extendedFamily = new Array("Moh", "Wamaitha")
myFamily = myFamily.concat(extendedFamily)
console.log("Parents are ", myFamily.slice(0,2))
console.log(myFamily)

/*how do i check if a number is a palindrome?*/
function isPalindrome(num){
    var str = num.toString()
    var reversedStr = str.split('').reverse().join('')
    return str === reversedStr
}
console.log(isPalindrome(9009))
console.log(isPalindrome(12341))

