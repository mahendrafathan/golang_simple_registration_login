const form = document.getElementById('form');
const mobileNumber = document.getElementById('mobileNumber');
const firstName = document.getElementById('firstName');
const lastName = document.getElementById('lastName');
const email = document.getElementById('email');
const gender = document.getElementById('gender');
const month = document.getElementById('month');
const date = document.getElementById('date');
const year = document.getElementById('year');

function showError(input, message) {
    const formControl = input.parentElement;
    formControl.className = 'form-control error';
    const small = formControl.querySelector('small');
    small.innerText = message;
    return true
}

function showSuccess(input, message) {
    const formControl = input.parentElement;
    formControl.className = 'form-control success';
    const small = formControl.querySelector('small');
    small.innerText = '';
}

var isIdnPhoneNumber = function(str) {
    return /^(^\+62\s?|^0)(\d{3,4}-?){2}\d{3,4}$/g.test(str)
}

function onlyNumberKey(evt) { 
    // Only ASCII charactar in that range allowed 
    var ASCIICode = (evt.which) ? evt.which : evt.keyCode 
    if (ASCIICode > 31 && (ASCIICode < 48 || ASCIICode > 57)) 
        return false; 
    return true; 
} 

function getRadioVal(name) {
    var val;
    // get list of radio buttons with specified name
    var radios = document.getElementsByName('gender')
    
    // loop through list of radio buttons
    for (var i=0, len=radios.length; i<len; i++) {
        if ( radios[i].checked ) { // radio checked?
            val = radios[i].value; // if so, hold its value in val
            break; // and break out of for loop
        }
    }

    return val; // return value of checked radio or undefined if none checked
}

// listener
form.addEventListener('submit', function(e){
    e.preventDefault();

    var isError = false
    if (mobileNumber.value === ''){
        isError = showError(mobileNumber, 'Mobile Number harus diisi')
    } else if (!isIdnPhoneNumber(mobileNumber.value)) {
        isError = showError(mobileNumber, 'Mobile Number salah')
    } else {
        showSuccess(mobileNumber)
    }

    if (firstName.value === ''){
        isError = showError(firstName, 'First Name harus diisi')
    } else {
        showSuccess(firstName)
    }

    if (lastName.value === ''){
        isError = showError(lastName, 'Last Name harus diisi')
    } else {
        showSuccess(lastName)
    }

    if (email.value === ''){
        isError = showError(email, 'Email harus diisi')
    } else {
        showSuccess(email)
    }

    var gender = getRadioVal('gender')
    var doB = year.value + '-' + month.value + '-' + date.value
    if (!isError){
        postRegister(parseInt(mobileNumber.value), 
        firstName.value, lastName.value, doB, gender, email.value)
    }
})

function postRegister(mobileNumber, firstName, lastName, doB, gender, email){
    $.ajax({
        url: `/register`,
        type: 'POST',
        contentType: 'application/json',
        data: JSON.stringify({    
            phone_number: mobileNumber,
            first_name: firstName,
            last_name: lastName,
            date_of_birth: doB,
            gender: gender,
            email: email,
        }), 
        success: function(msg){
            window.location = '/login';
        },
        error: function(XMLHttpRequest, textStatus, errorThrown) {
            alert("some error");
        }
    });
}

function login(email) {
    $.ajax({
        url: `/login`,
        type: 'POST',
        contentType: 'application/json',
        data: JSON.stringify({    
            email: email,
        }), 
        success: function(msg){
            window.location = '/';
        },
        error: function(XMLHttpRequest, textStatus, errorThrown) {
            alert("some error");
        }
    });
}