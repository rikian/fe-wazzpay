(async function () {
    // get access token and refesh token from local storage
    function ckLa() {
        const ls = localStorage.getItem("ls")
        if (!ls || ls == null) {
            return false
        }
        try {
            const lsObj = JSON.parse(ls)
            if (!lsObj.at || !lsObj.rt) {
                return false
            }
            return lsObj
        } catch (error) {
            return false
        }
    }

    // access token
    async function vAT(obj) {
        try {
            const body = JSON.stringify({
                "token": obj.at,
            })
            const request = await fetch("https://api.wazzpay.com/api/v1/user/auth/check-token", {
                method: "post",
                headers: {
                    "content-type": "application/json",
                    "content-length": body.length
                },
                body: body
            })
            const response = await request
            if (response.status !== 200) {
                return false
            }
            return true
        } catch (error) {
            return false
        }
    }

    const dtAc = ckLa()
    if (dtAc) {
        if (await vAT(dtAc)) {
            return window.location = "/"
        }
    }
})()