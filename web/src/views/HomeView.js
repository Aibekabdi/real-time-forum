import AbstractView from "./AbstractView.js";

export default class extends AbstractView{
    constructor(params) {
        super(params);
        this.setTitle("Home");
    }
    async init() {
        
    }
    async getHtml(){
        //home
        return `
        `;
    }

}