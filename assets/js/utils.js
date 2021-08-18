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
    let dloadDelim = "[download]"


    this.rows
    this.buffer = ""
    this.text = ""

    
    this.parse = function (row) {
        
        let text = ""
        let rows = this.buffer.split("\n")
        this.buffer += row

        rows.forEach(line => {
            if(line.includes(dloadDelim) && line.includes("%")){
                let chunks = line.split('[')
                line = ""
                for(let i = chunks.length -1; i != 0; i -= 1){
                    if(chunks[i].includes(" ETA ") || chunks[i].includes("100%")){
                        line = "["+ chunks[i]
                        break
                    }
                } 

            }
            text += line + "\n"
        })

        this.text = text
    }
}