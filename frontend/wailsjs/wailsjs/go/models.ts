export namespace main {
	
	export class Preset {
	    text: string;
	    min_ms: number;
	    max_ms: number;
	    loops: number;
	    loop_delay_ms: number;
	    loop_jitter_ms: number;
	
	    static createFrom(source: any = {}) {
	        return new Preset(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.text = source["text"];
	        this.min_ms = source["min_ms"];
	        this.max_ms = source["max_ms"];
	        this.loops = source["loops"];
	        this.loop_delay_ms = source["loop_delay_ms"];
	        this.loop_jitter_ms = source["loop_jitter_ms"];
	    }
	}

}

