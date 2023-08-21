function cal() {
    var number1 = document.getElementById("number1");
    var number2 = document.getElementById("number2");
    var result = document.getElementById("result");
    var s = parseInt(number1.value) + parseInt(number2.value);
    result.innerHTML = s;
}