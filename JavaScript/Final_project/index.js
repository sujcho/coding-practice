window.onload = function(){
    initialize();
}

function initialize(){
    validateInput()
}

function validateInput(){
    validateAmountOfLoan()
    validateAnnualInterest()
    validateRepaymentPeriod()
    validateZipcode()
}
function validateAmountOfLoan(){
    validate("amountOfLoan", /\d/)
}
function validateAnnualInterest(){
    validate("annualInterest", /\d/)
}
function validateRepaymentPeriod(){
    validate("repaymentPeriod", /\d/)
}
function validateZipcode(){
    validate("zipcode", /\d/)
}
function validate(element, testingForm){
    var textbox = document.getElementById(element);
    textbox.addEventListener("keypress", function(event){
        //alert('key pressed' + String.fromCharCode(event.charCode ) + count_length)
        var charCode = event.charCode;
        if(charCode !=0){
            if (!testingForm.test(String.fromCharCode(charCode))){
                event.preventDefault();
            }
        }
    })
}


/*
 * This script defines the calculate() function called by the event handlers
 * in HTML above. The function reads values from <input> elements, calculates
 * loan payment information, displays the results in <span> elements. It also
 * saves the user's data, displays links to lenders, and draws a chart.
 */
function calculate(){

    alert("onclick")

    var initial = document.getElementById("amountOfLoan").value;
    var annual_i = document.getElementsById("annualInterest").value;
    var period = document.getElementById("repaymentPeriod").value;

    //monthly interest
    var i = (annual_i * 0.01)/12;

    //total number of payment
    var n = period * 12

    var month_payment = initial * i * (Math.pow(1+ i, n)) / (Math.pow(1 + i, n) - 1);
    var total_payment = month_payment * n;
    var total_interest = total_payment - initial;


    document.getElementById("month_payment").innerHTML = ""+month_payment



}




