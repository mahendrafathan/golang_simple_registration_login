const form = document.getElementById('form');
const email = document.getElementById('email');

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

// listener
document.getElementById("login").addEventListener('click', function (e) {
    e.preventDefault();
    console.log('hallo');
    var isError = false
    if (email.value === '') {
        isError = showError(email, 'Email harus diisi')
    } else {
        showSuccess(email)
    }

    if (!isError) {
        login(email.value)
    }
    form.submit();
})

function login(email) {
    $.ajax({
        url: `/login`,
        type: 'POST',
        contentType: 'application/json',
        data: JSON.stringify({
            email: email,
        }),
        success: function (msg) {
            window.location = '/';
        },
        error: function (xhr, textStatus, errortdrown) {
            response = xhr.responseText;
            alert("got error " + response);
        }
    });
}

document.getElementById("register").addEventListener("click", function (e) {
    window.location = '/register';
});