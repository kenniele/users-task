let updateButton = document.getElementById("updateForm").querySelector("button")
console.log(updateButton)

if (updateButton) {
    updateButton.onclick = function (e) {
        e.preventDefault()
        let inputs = document.getElementById("updateForm").querySelectorAll("input")
        let block = document.getElementById("updateForm").querySelector("#block")
        let data = {}
        console.log(inputs)

        inputs.forEach(input => {
            data[input.name] = input.value
        })

        console.log(data)

        let xhr = new XMLHttpRequest()
        xhr.open("PUT", `/update/${data.passportNumber}`)
        xhr.setRequestHeader("Content-Type", "application/json")

        xhr.onload = function (e) {
            let response = JSON.parse(e.currentTarget.response)
            if (response && "Error" in response) {
                if (response.Error == null) {
                    block.innerText = "Аккаунт успешно обновлен"
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