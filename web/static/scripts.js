document.getElementById("myButton").addEventListener("click", function () {
    const fileInput = document.getElementById("fileInput");
    const file = fileInput.files[0];

    if (!file) {
        alert("Пожалуйста, выберите файл для загрузки.");
        return;
    }

    const formData = new FormData();
    formData.append("file", file);

    fetch("http://localhost:8080/api/save", {
        method: "POST", 
        body: formData,
    })
    .then(response => {
        if (!response.ok){
            throw new Error("Failed to get response");
        }
        return response.blob();
    })
    .then(data =>{
        const qrImage = URL.createObjectURL(data);
        document.getElementById("qrImage").src = qrImage;
    })
    .catch(error =>{
        console.error("Problem with your fetch operation:", error);
    })
});



