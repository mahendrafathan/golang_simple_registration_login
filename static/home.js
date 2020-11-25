const form = document.getElementById('form');

// listener
document.getElementById("logout").addEventListener("click", function () {
    console.log("hallo")
    $.ajax({
        url: `/logout`,
        type: 'GET',
        contentType: 'application/json',
        success: function (msg) {
            window.location = '/login';
        },
        error: function (xhr, textStatus, errortdrown) {
            response = xhr.responseText;
            alert("got error " + response);
        }
    });
    form.submit();
});

document.getElementById("snake").addEventListener("click", function () {
    console.log("hallo")
    $.ajax({
        url: `/snake`,
        type: 'GET',
        contentType: 'application/json',
        success: function (msg) {
            window.location = '/snake';
        },
        error: function (xhr, textStatus, errortdrown) {
            response = xhr.responseText;
            alert("got error " + response);
        }
    });
    form.submit();
});