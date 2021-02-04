(function($) {
    'use strict';
    $(function() {
        var ws;

        var ondisplay = function(e) {
            var responseMsg = JSON.parse(e.data)
            console.log("[display] parsed msg: ", responseMsg);
            console.log(`[display] command   : ${responseMsg.command}`);
            console.log(`[display] subCommand: ${responseMsg.subCommand}`);
            console.log(`[display] result    : ${responseMsg.result}`);
            console.log(`[display] resultMsg : ${responseMsg.resultMsg}`);
            console.log(`[display] datas      : ${responseMsg.data}`);

            var i = 0;
            responseMsg.data.forEach(element => {
                console.log(`[display] index           : (${i}), ${element.index}`);
                console.log(`[display] count           : (${i}), ${element.count}`);

                console.log(`[display] serialNo        : (${i}), ${element.serialNo}`);
                console.log(`[display] userImsi        : (${i}), ${element.userImsi}`);
                console.log(`[display] status          : (${i}), ${element.status}`);
                console.log(`[display] usedTime        : (${i}), ${element.usedTime}`);
                console.log(`[display] lastReceivedTime: (${i}), ${element.lastReceivedTime}`);
                console.log(`[display] watchDogOn      : (${i}), ${element.watchDogOn}`);

                $('#serialNo').html(element.serialNo);
                $('#userImsi').html(element.userImsi);
                $('#status').html(
                    `
                    <p>alterUser: ${element.status.alterUser}</p>
                    <p>lurUpRequest: ${element.status.lurUpRequest}</p>
                    <p>smsSend: ${element.status.smsSend}</p>
                    <p>call: ${element.status.call}</p>
                    <p>callDrop: ${element.status.callDrop}</p>
                    <p>pagingRequest: ${element.status.pagingRequest}</p>
                    <p>callRecvRequest: ${element.status.callRecvRequest}</p>
                    `
                );
                $('#usedTime').html(element.usedTime);
                $('#lastReceivedTime').html(element.lastReceivedTime);
                $('#watchDogOn').html((element.watchDogOn === false)? "false":"true");

                i++;
            });
        }

        var connect = function() {
            ws = new WebSocket("ws://" + window.location.host + "/monitor")
            ws.onopen = function(e) {
                console.log("[ws] onopen: ", arguments);
            }
            ws.onclose = function(e) {
                console.log("[ws] onclose: ", arguments);
            }
            ws.onmessage = function(e) {
                console.log("[ws] onmessage: ", arguments);
                ondisplay(e);
            }
        }
        connect();

        var webSocketInterval;
        $('#monitorOnceButton').on('click', function() {
            if(ws.readyState != ws.OPEN) {
                alert("disconnected webSocket")
            }
            ws.send(JSON.stringify({
                command: "all",
                subCommand: "once"
            }))
        });
        $('#monitorStartButton').on('click', function() {
            webSocketInterval = setInterval(function() {
                if(ws.readyState != ws.OPEN) {
                    alert("disconnected webSocket")
                    clearInterval(webSocketInterval);
                }
                ws.send(JSON.stringify({
                    command: "all",
                    subCommand: "start"
                }))
            }, 1000)
        })
        $('#monitorStopButton').on('click', function() {
            clearInterval(webSocketInterval);
        })

    });
})(jQuery);