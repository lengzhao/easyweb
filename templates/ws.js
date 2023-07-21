var sender;

function setSender(s) {
    sender = s;
}
function wsMessage(e) {
    // console.log("receive:", e)
    var obj = JSON.parse(e.data);
    switch (obj.type) {
        case "title":
            document.title = obj.msg
            break;
        case "js":
            $.getScript(obj.msg);
            break;
        case "css":
            $.ajax({
                url: obj.msg,
                dataType: 'text',
                success: () => {
                    $("head").append("<link>");
                    let css = $("head").children(":last");
                    css.attr({
                        rel: "stylesheet",
                        type: "text/css",
                        href: obj.msg
                    });
                }
            });
            break
        default:
            var element = $("#" + obj.id);
            if (element.length > 0) {
                if (obj.msg == 0) {
                    $("#" + obj.id).remove();
                    // console.log("删除html")
                    return
                }
                if (obj.type == "label") {
                    $("#" + obj.id).text(obj.msg)
                } else {
                    $("#" + obj.id).html(obj.msg)
                }

                // console.log("修改html")
            } else {
                g = document.createElement('div');
                g.setAttribute("id", obj.id);

                if (obj.type == "label") {
                    g.innerText = obj.msg;
                } else {
                    g.innerHTML = obj.msg;
                }

                $("div:first").append(g);
                // $("#" + obj.id).html(obj.msg);
                // console.log("添加HTML")
            }
            bindEvent(obj.type, obj.id);
            break;
    }
};

function bindEvent(typ, id) {
    switch (typ) {
        case "input":
            $("#" + id + " input").on("input", function () {
                var info = {
                    type: typ,
                    msg: $(this).val(),
                    id: id
                }
                sender.send(JSON.stringify(info));
            });
            break;
        case "button":
            $("#" + id + " button").click(function () {
                var info = {
                    type: "button",
                    msg: "click",
                    id: id
                }
                sender.send(JSON.stringify(info));
            });
            break
        case "form":
            $("#" + id + " form").submit(function (event) {
                event.preventDefault();
                const data = new FormData(event.target);
                const value = Object.fromEntries(data.entries());
                console.log("form:",value)
                var info = {
                    type: "form",
                    msg: JSON.stringify(value),
                    id: id
                }
                // $(this).serialize(),
                sender.send(JSON.stringify(info));
                return false;
            });
            break

        default:
            break;
    }
}
