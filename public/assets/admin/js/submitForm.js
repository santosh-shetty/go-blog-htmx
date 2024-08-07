var baseUrl = "http://localhost:9000";
// Display a success message after the AJAX call completes
const displayMessage = (message, delay) => {
    $.notify(`<i class="ti-plus btn-icon-append"></i><strong>${message}</strong>`, {
        allow_dismiss: true,
        delay: delay,
        showProgressbar: true,
        timer: 300
    });
}
// =============== Start Blog Script ==================== 
//  Add Blog Form
$("#addBlog").submit(function (e) {

    e.preventDefault();
    $("#submtBtn").prop("disabled", true)

    var formData = new FormData(this);
    var description = tinyMCE.get('description').getContent()
    formData.set("description", description)
    $.ajax({
        type: "POST",
        url: `${baseUrl}/admin/blog/add`,
        contentType: false,
        processData: false,
        data: formData,
        success: function (data) {
            if (data.status === "success") {
                displayMessage("Blog Added Successfully!", 2000);
                setInterval(() => {
                    $("#submtBtn").prop("disabled", false)
                    window.location.href = `${baseUrl}/admin/blog/list`
                }, 2000);
            } else {
                $("#submtBtn").prop("disabled", false)
                console.log('An error occurred:', data.message);
            }

        },
        error: function (data) {
            console.log('An error occurred.');
            console.log(data);
            $("#submtBtn").prop("disabled", false)

        },
    });

});
//  Update Blog Form
$("#updateBlog").submit(function (e) {
    e.preventDefault();
    var id = $('#blogId').val();
    $("#submtBtn").prop("disabled", true)
    var formData = new FormData(this);
    var description = tinyMCE.get('description').getContent()
    formData.set("description", description)

    $.ajax({
        type: "POST",
        url: `${baseUrl}/admin/blog/update/${id}`,
        contentType: false,
        processData: false,
        data: formData,
        success: function (data) {
            if (data.status === "success") {
                displayMessage("Blog Updated Successfully!", 2000);
                setTimeout(() => {
                    window.location.href = `${baseUrl}/admin/blog/list`;
                }, 2000);
            } else {
                $("#submtBtn").prop("disabled", false)
                console.log('An error occurred:', data.message);
            }
        },
        error: function (data) {
            $("#submtBtn").prop("disabled", false)
            console.log('An error occurred.');
            console.log(data);
        },
    });

});
// =============== End  Blog Script ==================== 




// =============== Start Category Script ==================== 

//  Add Category Form
$("#addCategory").submit(function (e) {

    e.preventDefault();
    $("#submtBtn").prop("disabled", true)
    var formData = new FormData(this);

    $.ajax({
        type: "POST",
        url: `${baseUrl}/admin/category/add`,
        contentType: false,
        processData: false,
        data: formData,
        success: function (data) {
            if (data.status === "success") {
                displayMessage("Category Added Successfully!", 2000);
                setTimeout(() => {
                    window.location.href = `${baseUrl}/admin/category/list`;
                }, 2000);
            } else {
                $("#submtBtn").prop("disabled", false)
                console.log('An error occurred:', data.message);
            }
        },
        error: function (data) {
            $("#submtBtn").prop("disabled", false)
            console.log('An error occurred.');
            console.log(data);
        },
    });

});
// Update Category
$("#updateCategory").submit(function (e) {

    e.preventDefault();
    $("#submtBtn").prop("disabled", true)
    var formData = new FormData(this);
    var id = $('#categoryId').val();
    $.ajax({
        type: "POST",
        url: `${baseUrl}/admin/category/update/${id}`,
        contentType: false,
        processData: false,
        data: formData,
        success: function (data) {
            if (data.status === "success") {
                displayMessage("Category Updated Successfully!", 2000);
                setTimeout(() => {
                    window.location.href = `${baseUrl}/admin/category/list`;
                }, 2000);
            } else {
                $("#submtBtn").prop("disabled", false)
                console.log('An error occurred:', data.message);
            }
        },
        error: function (data) {
            $("#submtBtn").prop("disabled", false)
            console.log('An error occurred.');
            console.log(data);
        },
    });

});