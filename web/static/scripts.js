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

document.getElementById("secondButton").addEventListener("click", function (){
    fetch("http://localhost:8080/api/img/01930ad7-34ba-7cb0-b0e2-94637e965cc8?extension=jpeg").then(response => {
        if (!response.ok) {
            throw new Error("Network response was not ok " + response.statusText);
        }
        return response.blob();
    }).then(data => {
        // Обработка полученных данных (например, добавление изображения на страницу)
        const imageUrl = URL.createObjectURL(data);
        document.getElementById("qrImage").src = imageUrl; // Предполагается, что у вас есть элемент <img id="myImage">
    })
        .catch(error => {
            console.error("There was a problem with your fetch operation:", error);
        });
})

