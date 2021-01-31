var app = new Vue({
    el: "#app",
    data: {
        value1: '',
        value2: '',
        result: ''
    },
    methods: {
        add: function () {
            this.result = add(this.value1, this.value2)
        },
        subtract: function () {
            this.result = subtract(this.value1, this.value2)
        }
    }
})
