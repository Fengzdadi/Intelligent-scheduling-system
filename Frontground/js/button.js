window.onload = function () {
    var cells = document.getElementById('monitor').getElementsByTagName('td');
    var clen = cells.length;
    var currentFirstDate;
    var formatDate = function (date) {
        var year = date.getFullYear() + '年';
        var month = (date.getMonth() + 1) + '月';
        var day = date.getDate() + '日';
        var week = ['周日', '周一', '周二', '周三', '周四', '周五', '周六'][date.getDay()]+"<br/>";
        return week + month + day;
    };
    var addDate = function (date, n) {
        date.setDate(date.getDate() + n);
        return date;
    };
    var setDate = function (date) {
        var week = date.getDay() - 1;
        date = addDate(date, week * -1);
        currentFirstDate = new Date(date);
        for (var i = 0; i < clen; i++) {
            cells[i].innerHTML = formatDate(i == 0 ? date : addDate(date, 1));
        }
    };
    
    document.getElementById('last-week').onclick = function () {
        setDate(addDate(currentFirstDate, -7));
    };
    document.getElementById('next-week').onclick = function () {
        setDate(addDate(currentFirstDate, 7));
    };
    setDate(new Date());
}