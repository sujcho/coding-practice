window.onload = function(){
    addValidation();
}

function addValidation(){
    validateAmountOfLoan()
    validateAnnualInterest()
    validateRepaymentPeriod()
    validateZipcode()
}
function validateAmountOfLoan(){
    validate("amountOfLoan", /\d/ , "msg1", function(event){
        if(event.keyCode == 9){
            if( this.value.length > 10){
                var msgbox = document.getElementById("msg1");
                msgbox.innerHTML="Number is too big";
            }
        }
    })
}
function validateAnnualInterest(){
    validate("annualInterest", /\d/, "msg2", function(event){
        if(event.keyCode == 9){
            if( this.value > 100){
                var msgbox = document.getElementById("msg2");
                msgbox.innerHTML="Please type the correct interest";
            }
        }
    })
}
function validateRepaymentPeriod(){
    validate("repaymentPeriod", /\d/, "msg3",function(event){
        if(event.keyCode == 9){
            if( this.value > 50){
                var msgbox = document.getElementById("msg3");
                msgbox.innerHTML="Please type the correct period: max.50";
            }
        }
    })
}
function validateZipcode(){
    validate("zipcode", /\d/, "msg4", function(event){
        if(event.keyCode == 9){
            if( this.value.length > 5){
                var msgbox = document.getElementById("msg4");
                msgbox.innerHTML="Please type the correct zip code: 5 digit";
            }
        }
    })
}
function validate(element, testingForm ,msg, moreCheck){
    var textbox = document.getElementById(element);
    textbox.addEventListener("keypress", function(event){

        //alert('key pressed' + String.fromCharCode(event.charCode ) + count_length)
        var charCode = event.charCode;
        var msgbox = document.getElementById(msg);
        msgbox.innerHTML="";
        var msgbox2 = document.getElementById("msg5");
        msgbox2.innerHTML = ""
        if(charCode !=0){
            if (!testingForm.test(String.fromCharCode(charCode))){
                event.preventDefault();
                msgbox.innerHTML = "Please enter the number."
            }
        }
    })
    textbox.addEventListener("keydown", moreCheck);
}

/*
 * This script defines the calculate() function called by the event handlers
 * in HTML above. The function reads values from <input> elements, calculates
 * loan payment information, displays the results in <span> elements. It also
 * saves the user's data, displays links to lenders, and draws a chart.
 */
function calculate(){

    var initial = document.getElementById("amountOfLoan").value;
    var annual_i = document.getElementById("annualInterest").value;
    var period = document.getElementById("repaymentPeriod").value;

    //if anything is null, message error
    if(!initial||!annual_i||!period){
        var msgbox = document.getElementById("msg5");
        msgbox.innerHTML = "Please type all values";
        return
    }

    //monthly interest
    var i = (annual_i * 0.01)/12;

    //total number of payment
    var n = period * 12

    var month_payment = Math.round(initial * i * (Math.pow(1+ i, n)) / (Math.pow(1 + i, n) - 1));
    var total_payment = Math.round(month_payment * n);
    var total_interest = Math.round(total_payment - initial);

    document.getElementById("month_payment").innerHTML = ""+month_payment
    document.getElementById("total_payment").innerHTML = total_payment
    document.getElementById("total_interest").innerHTML = total_interest

    getLenders(initial, annual_i,period)
    chart(initial, total_interest, month_payment, total_payment)

}


function getLenders(amount, apr, years, zipcode){

    var user_data ={}
    user_data.amount = amount
    user_data.apr = apr
    user_data.years = years

    var user_json = JSON.stringify(user_data);

    var request;
    if (window.XMLHttpRequest) {
        request = new XMLHttpRequest();
    } else {
        request = new ActiveXObject("Microsoft.XMLHTTP");
    }
    request.open('GET', 'data.json');
    request.onreadystatechange = function() {
        if ((request.readyState===4) && (request.status===200)) {
            var items = JSON.parse(request.responseText);
            var output = '<ul>';
            for (var key in items) {
                output += '<li>' + items[key].name + '<a href=' + items[key].link + '>' + items[key].link + '</a>' +'</li>';
            }
            output += '</ul>';
            document.getElementById('landers').innerHTML = output;
        }
    };
    request.send();
}

function chart(principal, interest, monthly, payments) {
    var canvas = document.getElementById('chart');
    if (canvas.getContext) {
        var ctx = canvas.getContext('2d');

        //ctx.fillRect(x, y, width, height);
        //normalize
        //principal
        ctx.fillStyle = 'green';
        ctx.fillRect(0, 100, 70, 200);

        height = (interest/principal) * 200;
        ctx.fillStyle = 'blue';
        ctx.fillRect(80, 100, 70, 100);


        height = (monthly/principal) * 200;
        ctx.fillStyle = 'red';
        ctx.fillRect(160, 100, 70, height);

        height = (payments/principal) * 200;
        ctx.fillStyle = 'yellow';
        ctx.fillRect(240, 100, 70, height);


    }
}
