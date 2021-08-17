const audiovideo = 2
const audioonly = 1
const videoonly = 0 

// return formats with type field
function orderAndSetType(formats) {
    let audioVideoFormats = []
    let audioOnlyFormats = []
    let videoOnlyFormats = []

    formats.forEach(format => {
        if(format.acodec == 'none'){
            format.type = videoonly
            videoOnlyFormats.push(format)
        } else if(format.vcodec == 'none'){
            format.type = audioonly
            audioOnlyFormats.push(format)
        } else{
            format.type = audiovideo
            audioVideoFormats.push(format)
        }
    });
    return [].concat(audioVideoFormats, audioOnlyFormats, videoOnlyFormats)
}


function Parser(){
    this.destination
    this.warn
    this.progress = 0
    
    this.parse = function (row) {
    
        if(0){}
        if(!this.destination){
            let dest = row.match(/\[download\] Destination: (.*)/)
            if(dest){
                this.destination = dest[1]
            }
        }
    }
}