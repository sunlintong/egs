/**
 * Created by gali on 2019/1/30.
 */
$(function (){
    /**登录切换*/
    $("#doctor").addClass('selected');
    $("#heading-login").hide();
    $("#doctor-login").show();

    $("#doctor").click(function(){
        $("#heading").removeClass('selected');
        $("#doctor").addClass('selected');
        $("#heading-login").hide();
        $("#doctor-login").show();
    });
    $("#heading").click(function(){
        $("#heading").addClass('selected');
        $("#doctor").removeClass('selected');
        $("#heading-login").show();
        $("#doctor-login").hide();
    });


    /**通过enter实现登录*/
    $(document).keyup(function(e){
        if(e.keyCode == 13){
            //判断是哪一个角色登录
            if(($("#doctor-login").is(':visible'))&&($("#heading-login").is(":hidden"))){
                $("#doctor-reg").trigger('click')
            }else {
                $("#heading-reg").trigger('click')
            }
        }
    });

    $("#doctor-reg").click(function(){
        var username = $("#username1").val();
        var password = $("#password1").val();
        if(username == ""||password ==""){
            alert("用户名,密码不能为空");
        }
        //console.log(username,password)
        $.post('/doctor/login',{user:username, pwd:password},function(e){
            if(e.dode == 0){
                console.log("登录成功")
            }else {
                console.log("用户名或密码错误")
            }
        }).error(function(){
            console.log("网络异常")
        })
    })

    $("#heading-reg").click(function(){
        var username = $("#username2").val();
        var password = $("#password2").val();
        console.log(username,password)
    })
})
function check_U_P(username, password){
    $.ajax({
        type: "POST",
        dataType: "json",
        url: "/login",
        contentType: "application/json",
        data: JSON.stringify({
            username: username,
            password: password
        }),
        success: function(data){

        }
    })
}
