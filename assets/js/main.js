const THUMBNAIL_DEFAULT_URL = "https://via.placeholder.com/336x188.png?text=Image+not+found"

new Vue({
    el: "#main",
    data: {
        url: "",
        inProgress: false,
        video: {}
    },
    methods: {
        "getInfo": function(e){
            if(this.url.length < 8 || this.inProgress){
                return
            }

            this.video = {}

            this.inProgress = true
            fetch("/api/get-info", {
                method: 'POST',
                body: this.url.trim()
            })
            .then(resp => resp.text())
            .then(res => {

                this.video =  JSON.parse(res)
                //console.log(this.video)
            })
            .finally(() => this.inProgress = false)
        },
        "download": function (id) {
            if(this.inProgress){
                return
            }

            this.inProgress = true

            let total = 0

            let stream = new XMLHttpRequest()

            stream.open("GET", "/api/download/best", true)
            
            stream.onprogress = function (event) {
                if(event.loaded){
                    this.inProgress = false
                    return
                }
                total = event.total 
                console.log("PROGRESS:", stream.responseText)
            }

  

            stream.send()

        }
    },
    computed: {
        videoDuration() {
            d = this.video.duration 

            res = ""
	        spr = ":"

	        hours = parseInt(d / 3600)
	        rems = d % 3600
	        minutes = parseInt(rems / 60)
	        seconds = d % 60

            if( hours > 0 ){
                res +=  hours + spr
            }

            res += (minutes < 10 ? "0" : "" ) + minutes + spr
            res += (seconds < 10 ? "0" : "" ) + seconds

            return res

        },
        videoThumbnail(){
            if(this.video.extractor_key == "Youtube" && this.video.thumbnails){
                let maxIndex = 0
                let maxW = 0


                this.video.thumbnails.forEach((el, i) => {
                    if(el.width > maxW){
                        maxW = el.width
                        maxIndex = i
                    }
                });
                console.log(this.video.thumbnails[maxIndex])
                return this.video.thumbnails[maxIndex].url
            }
            return this.video.thumbnail || THUMBNAIL_DEFAULT_URL
        }
    }
})
