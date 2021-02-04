(function($) {
    'use strict';
    $(function() {
        
        $('#submit').on('click', function(event) {
            event.preventDefault();

            var id = $('#id').val();
            var password = $('#password').val();

            console.log(id, password);
            console.log(id.length, password.length);
            if(id.length > 0 && password.length) {
                console.log("login go");
                $.post('/login', {id: id, password: password})
                .done(function(msg) {
                    console.log("done: ", msg)
                    if(msg.result) {
                        var monitorUri = "http://" + window.location.host + "/monitor.html";
                        $(location).attr('href', monitorUri);
                    } else {
                        alert("login fail")
                    }
                })
                .fail(function(xhr, status, error) {
                    alert("login fail")
                });
            }
        });

    });
})(jQuery);