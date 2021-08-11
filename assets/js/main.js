new Vue({
    el: "#main",
    data: {
        url: "",
        inProgress: false,
        video: {}
    },
    methods: {
        "getInfo": function(){
            if(this.url.length < 8){
               console.log('text is short')
                return
            }
            
            this.inProgress = true
            fetch("/api/get-info", {
                method: 'POST',
                body: this.url.trim()
            })
            .then(resp => resp.text())
            .then(res => {
                this.video =  JSON.parse(res)
            })
            .finally(() => {this.inProgress = false})
        }
    }
})
