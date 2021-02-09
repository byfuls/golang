(function($) {
    'use strict';
    $(function() {
        var ws;

		var insertChannelList = function(index, element) {
			var channelIndex = `<td>${index}</td>`;
			var serialNo = `<td class="serialNo-${index}">${element.serialNo}</td>`;
			var userImsi = `<td class="userImsi-${index}">${element.userImsi}</td>`;
			var status_alterUser = (element.status.alterUser == false)?
				`<tr><td class="status-td">Alter User</td><td class="status-td status-alterUser-${index}">○</td></tr>`:
				`<tr><td class="status-td">Alter User</td><td class="status-td status-alterUser-${index}">●</td></tr>`;
			var status_lurUpRequest = (element.status.lurUpRequest == false)?
				`<tr><td class="status-td">Lur Up Request</td><td class="status-td status-lurUpRequest-${index}">○</td></tr>`:
				`<tr><td class="status-td">Lur Up Request</td><td class="status-td status-lurUpRequest-${index}">●</td></tr>`;
			var status_smsSend = (element.status.smsSend == false)?
				`<tr><td class="status-td">Sms Send</td><td class="status-td status-smsSend-${index}">○</td></tr>`:
				`<tr><td class="status-td">Sms Send</td><td class="status-td status-smsSend-${index}">●</td></tr>`;
			var status_call = (element.status.call == false)?
				`<tr><td class="status-td">Call</td><td class="status-td status-call-${index}">○</td></tr>`:
				`<tr><td class="status-td">Call</td><td class="status-td status-call-${index}">●</td></tr>`;
			var status_callDrop = (element.status.callDrop == false)?
				`<tr><td class="status-td">Call Drop</td><td class="status-td status-callDrop-${index}">○</td></tr>`:
				`<tr><td class="status-td">Call Drop</td><td class="status-td status-callDrop-${index}">●</td></tr>`;
			var status_pagingRequest = (element.status.pagingRequest == false)?
				`<tr><td class="status-td">Paging Request</td><td class="status-td status-pagingRequest-${index}">○</td></tr>`:
				`<tr><td class="status-td">Paging Request</td><td class="status-td status-pagingRequest-${index}">●</td></tr>`;
			var status_callRecvRequest = (element.status.callRecvRequest == false)?
				`<tr><td class="status-td">Call Recv Request</td><td class="status-td status-callRecvRequest-${index}">○</td></tr>`:
				`<tr><td class="status-td">Call Recv Request</td><td class="status-td status-callRecvRequest-${index}">●</td></tr>`;
			var status = "<td><table class='status-table'><tbody>"
						+ status_alterUser
						+ status_lurUpRequest
						+ status_smsSend
						+ status_call
						+ status_callDrop
						+ status_pagingRequest
						+ status_callRecvRequest
						+ "</tbody></table></td>";
			var usedTime = `<td class="usedTime-${index}">${element.usedTime}</td>`;
			var lastReceived = `<td class="lastReceived-${index}">${element.lastReceivedTime}</td>`;
			var watchDogOn = (element.watchDogOn === false)?
								`<td class="watchDogOn-${index}">F</td>`:
								`<td class="watchDogOn-${index}">T</td>`;

			return "<tr>" + channelIndex + serialNo + userImsi + status + usedTime + lastReceived + watchDogOn + "</tr>";
		}

		var i = 0;
        var ondisplay = function(e) {
            var responseMsg = JSON.parse(e.data)
            //console.log("[display] parsed msg: ", responseMsg);
            //console.log(`[display] command   : ${responseMsg.command}`);
            //console.log(`[display] subCommand: ${responseMsg.subCommand}`);
            //console.log(`[display] result    : ${responseMsg.result}`);
            //console.log(`[display] resultMsg : ${responseMsg.resultMsg}`);
            //console.log(`[display] datas      : ${responseMsg.data}`);

			var j = 0;
            responseMsg.data.forEach(element => {
                //console.log(`[display] index           : (${i}), ${element.index}`);
                //console.log(`[display] count           : (${i}), ${element.count}`);

                //console.log(`[display] serialNo        : (${i}), ${element.serialNo}`);
                //console.log(`[display] userImsi        : (${i}), ${element.userImsi}`);
                //console.log(`[display] status          : (${i}), ${element.status}`);
                //console.log(`[display] usedTime        : (${i}), ${element.usedTime}`);
                //console.log(`[display] lastReceivedTime: (${i}), ${element.lastReceivedTime}`);
                //console.log(`[display] watchDogOn      : (${i}), ${element.watchDogOn}`);

				/* insert channel list */
				j++;
				//console.log(`j: ${j} i: ${i}`);
				if(j == i) {
					$('.serialNo-'+i).text(element.serialNo);
					$('.userImsi-'+i).text(element.userImsi);
					(element.status.alterUser === false)?
							$('.status-alterUser-'+i).text('○'):$('.status-alterUser-'+i).text('●');
					(element.status.lurUpRequest === false)?
							$('.status-lurUpRequest-'+i).text('○'):$('.status-lurUpRequest-'+i).text('●');
					(element.status.smsSend === false)?
							$('.status-smsSend-'+i).text('○'):$('.status-smsSend-'+i).text('●');
					(element.status.call === false)?
							$('.status-call-'+i).text('○'):$('.status-call-'+i).text('●');
					(element.status.callDrop === false)?
							$('.status-callDrop-'+i).text('○'):$('.status-callDrop-'+i).text('●');
					(element.status.pagingRequest === false)?
							$('.status-pagingRequest-'+i).text('○'):$('.status-pagingRequest-'+i).text('●');
					(element.status.callRecvRequest === false)?
							$('.status-callRecvRequest-'+i).text('○'):$('.status-callRecvRequest-'+i).text('●');
					(element.status.usedTime=== false)?
							$('.status-usedTime-'+i).text('○'):$('.status-usedTime-'+i).text('●');
					(element.status.lastReceived=== false)?
							$('.status-lastReceived-'+i).text('○'):$('.status-lastReceived-'+i).text('●');
					(element.watchDogOn === false)?
						$('.watchDogOn-'+i).text("F"):$('.watchDogOn-'+i).text("T");
				} else {
                	i++;
					var channelListHtml = insertChannelList(i, element);
					$('.channelList').append(`${channelListHtml}`);
				}
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
                //console.log("[ws] onmessage: ", arguments);
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
            }, 500)
        })
        $('#monitorStopButton').on('click', function() {
            clearInterval(webSocketInterval);
        })

    });
})(jQuery);
