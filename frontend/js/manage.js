const endTaskContainer = document.getElementById("currentTasks")

endTaskContainer.addEventListener("click", (e) => {
    if (e.target.tagName === "BUTTON") {
        let userPassport = e.target.getAttribute("data-passport")
        let TaskID = e.target.getAttribute("data-task")
        let xhr = new XMLHttpRequest()
        console.log(userPassport, TaskID, e)
        let block = endTaskContainer.querySelector("#blockDelete")
        xhr.open("POST", `/man/end/${userPassport}/${TaskID}`)
        xhr.setRequestHeader("Content-Type", "application/json")
        xhr.onload = function (e) {
            let response = JSON.parse(e.currentTarget.response)
            if (response && "Error" in response) {
                let old = block
                block = `<div class="alert alert-danger" role="alert">Ошибка: ${response.Error}</div>`
                setTimeout(() => {
                    block = old
                }, 1000)
            }
        }
        xhr.send({
            "Passport": userPassport,
            "TaskID": TaskID,
        })
    }
})

const createTaskContainer = document.getElementById("createModal")

createTaskContainer.addEventListener("click", (e) => {
    if (e.target.tagName === "BUTTON") {
        let userPassport = e.target.getAttribute("data-passport")
        let inputs = createTaskContainer.querySelectorAll("input")
        let data = {
            "Passport": userPassport,
        }
        inputs.forEach(input => {
            data[input.name] = input.value
        })
        let xhr = new XMLHttpRequest()
        let block = createTaskContainer.querySelector("#blockCreate")
        console.log(userPassport)
        xhr.open("POST", `/man/begin/${userPassport}/${data.TaskName}`)
        xhr.setRequestHeader("Content-Type", "application/json")
        xhr.onload = function (e) {
            let response = JSON.parse(e.currentTarget.response)
            if (response && "Error" in response) {
                let old = block
                block = `<div class="alert alert-danger" role="alert">Ошибка: ${response.Error}</div>`
                setTimeout(() => {
                    block = old
                }, 1000)
            }
        }
        xhr.send(JSON.stringify(data))
    }
})