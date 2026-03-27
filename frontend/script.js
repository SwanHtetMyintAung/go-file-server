const form = document.querySelector("form");

const input = form.querySelector("input");
const btn = form.querySelector("button");

btn.addEventListener("click",(e)=>{
    const form = new FormData;
    form.append("file",input.files[0])
    console.log(form)
    e.preventDefault();
    fetch("http://localhost:3000/file",{
        method:"POST",
        headers:{
            "Content-Type": "text/plain"
        },
        body:form
    

    })
    .then(res => res.json())
    
})