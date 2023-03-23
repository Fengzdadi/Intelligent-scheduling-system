new Vue({
    el: '#app',
    data: {
        time: {
            y: 2023,//年
            M: 01,//月
            day: 12,//日
            h: 13,//时
            m: 18,//分
            s: 00,//秒
            week: '星期四'
        }
    },
    methods: {
        clock: function () {
            let oDate = new Date();
            this.time = {
                y: oDate.getFullYear(),
                M: oDate.getMonth() + 1,
                day: oDate.getDate(),
                h: oDate.getHours(),
                m: oDate.getMinutes(),
                s: oDate.getSeconds(),
                week: oDate.getDay(),
            }
            //月数据处理，个位数时补0
            if (parseInt(this.time.M) < 10) {
                this.time.M = '0' + parseInt(this.time.M)
            }
            //日数据处理，个位数时补0
            if (parseInt(this.time.day) < 10) {
                this.time.day = '0' + parseInt(this.time.day)
            }
            //时数据处理，个位数时补0
            if (parseInt(this.time.h) < 10) {
                this.time.h = '0' + parseInt(this.time.h)
            }
            //分数据处理，个位数时补0
            if (parseInt(this.time.m) < 10) {
                this.time.m = '0' + parseInt(this.time.m)
            }
            //秒数据处理，个位数时补0
            if (parseInt(this.time.s) < 10) {
                this.time.s = '0' + parseInt(this.time.s)
            }
            //星期处理，展示星期几
            let week = parseInt(this.time.week)
            switch (week) {
                case 1: this.time.week = '星期一'; break;
                case 2: this.time.week = '星期二'; break;
                case 3: this.time.week = '星期三'; break;
                case 4: this.time.week = '星期四'; break;
                case 5: this.time.week = '星期五'; break;
                case 6: this.time.week = '星期六'; break;
                case 7: this.time.week = '星期日'; break;
            }
        }
    },
    beforeCreate: function () {
        console.log('创建前');
    },
    created: function () {
        console.log('创建后');
    },
    beforeMount: function () {
        console.log('挂载前');
    },
    mounted: function () {
        console.log('挂载后');
        setInterval(this.clock, 1000);//每隔一秒执行一次
        this.clock();//页面加载时候执行一次
    },
    beforeUpdate: function () {
        console.log('数据更新前');
    },
    updated: function () {
        console.log('数据更新后');
    },
    beforeDestroy: function () {
        console.log('销毁前');
    },
    destroyed: function () {
        console.log('销毁后');
    },
})