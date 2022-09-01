import AbstractView from "./AbstractView.js";

export default class extends AbstractView{
    constructor() {
        super();
        this.setTitle("Signin");
    }
    async getHtml(){
        //signin
        return `
        `;
    }

}