const form = document.getElementById('form');
const mobileNumber = document.getElementById('mobileNumber');
const firstName = document.getElementById('firstName');
const lastName = document.getElementById('lastName');
const email = document.getElementById('email');
const gender = document.getElementById('gender');
const month = document.getElementById('month');
const date = document.getElementById('date');
const year = document.getElementById('year');

// listener
form.addEventListener('submit', function(e){
    e.preventDefault();
    $.ajax({
        url: `/logout`,
        type: 'GET',
        contentType: 'application/json',
        success: function(msg){
            window.location = '/login';
        },
        error: function (xhr, textStatus, errortdrown) {
            response = xhr.responseText;
            alert("got error " + response);
        }
    });
})
