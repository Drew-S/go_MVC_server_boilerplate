'use strict';

(() => {

    const template = document.createElement("template");
    template.innerHTML = /* html */`
        <button class="btn btn-primary">TEXT</button>
    `;

    window.customElements.define('my-thing', class extends HTMLElement {
        constructor() {
            super();

            // const shadow = this.attachShadow({ mode: 'open'});

            // shadow.appendChild(template.content.cloneNode(true));
            this.appendChild(template.content.cloneNode(true));
        }
    });
    
})();