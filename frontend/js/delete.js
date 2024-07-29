let deleteButton = document.getElementById("deleteForm").querySelector("button")

if (deleteButton) {
    deleteButton.onclick = function (e) {
        e.preventDefault()
        let inputs = document.getElementById("deleteForm").querySelectorAll("input")
        let block = document.getElementById("deleteForm").querySelector("#block")
        let data = {}
        console.log(inputs)

        inputs.forEach(input => {
            data[input.name] = input.value
        })

        console.log(data)

        let xhr = new XMLHttpRequest()
        xhr.open("DELETE", `/delete/${data.passportNumber}`)
        xhr.setRequestHeader("Content-Type", "application/json")

        xhr.onload = function (e) {
            let response = JSON.parse(e.currentTarget.response)
            if (response && "Error" in response) {
                if (response.Error == null) {
                    block.innerText = "Аккаунт успешно удален"
                    setTimeout(() => {
                        block.innerText = ""
                    }, 1000)
                } else {
                    block.innerText = `Возникла ошибка ${response.Error}`
                    setTimeout(() => {
                        block.innerText = ""
                    }, 1000)
                }
            } else {
                block.innerText = "Некорректные данные"
                setTimeout(() => {
                    block.innerText = ""
                }, 1000)
            }
        }
        xhr.send(JSON.stringify(data))
    }
}