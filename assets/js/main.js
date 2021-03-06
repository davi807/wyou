const THUMBNAIL_DEFAULT_URL = "https://via.placeholder.com/336x188.png?text=Image+not+found"

new Vue({
    el: "#main",
    data: {
        url: "",
        inProgress: false,
        filters: {av: true, ao: true, vo: true},
        downloading: false,
        progressText: "",
        video: {}
    },
    methods: {
        "getInfo": function(e){
            if(this.url.length < 8 || this.inProgress){
                return
            }

            this.reset()

            this.inProgress = true
            fetch("/api/get-info", {
                method: 'POST',
                body: this.url.trim()
            })
            .then(resp => resp.text())
            .then(res => {
                let video =  JSON.parse(res)
                if(video.formats){
                    video.formats = video.formats.sort((a, b) => b.filesize - a.filesize)
                    video.formats = orderAndSetType(video.formats)
                }
                this.video = video
            })
            .finally(() => this.inProgress = false)
        },
        "download": function (id) {
            if(this.inProgress){
                return
            }

            this.inProgress = true

            let self = this 
            let total = 0
            let parser = new Parser()
            let stream = new XMLHttpRequest()

            stream.open("GET", "/api/download/"+id, true)

            self.progressText = ""
            self.downloading = true

            stream.onprogress = function () {
                let responseText = stream.responseText.substring(total) 
                total = stream.responseText.length 

                parser.parse(responseText)
                self.$refs.progressText.innerText = parser.text

                if(responseText.includes("##DONE##")){
                    self.inProgress = false
                }
            }

            stream.send()
            window.scroll(0, 0);  
        },
        "updateFormats": function(){
            if(!this.video.formatsBackup){
                this.video.formatsBackup = this.video.formats
            }

            this.video.formats = []

            this.video.formatsBackup.forEach(format => {
                if( (format.type === audiovideo && this.filters.av) || 
                    (format.type === audioonly && this.filters.ao) ||
                    (format.type === videoonly && this.filters.vo)
                ){
                    this.video.formats.push(format)
                }
            })

        },
        "reset": function(){
            this.filters = {av: true, ao: true, vo: true},
            this.downloading = false,
            this.progressText = "",
            this.video = {}
        },
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
                return this.video.thumbnails[maxIndex].url
            }
            return this.video.thumbnail || THUMBNAIL_DEFAULT_URL
        }
    }
})

function updateDL(){
    document.querySelector("#update-btn").remove()
    let container = document.querySelector("#update-info")
    container.innerText = "Updating, please wait..."

    let stream = new XMLHttpRequest()

    let done

    stream.onprogress = function (event) {
        if(done){
            return
        }
        if(stream.responseText.includes("##DONE##")){
            container.innerText = "Update finished!"
            done = true
            alert(stream.responseText)
        }
    }

    stream.open("GET", "/api/update/", true)
    stream.send()

}