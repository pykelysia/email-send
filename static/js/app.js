// 邮件发送 WebUI 脚本
document.addEventListener('DOMContentLoaded', function() {
    const form = document.getElementById('emailForm');
    const submitBtn = document.getElementById('submitBtn');
    const btnText = submitBtn.querySelector('.btn-text');
    const btnLoading = submitBtn.querySelector('.btn-loading');
    const resultDiv = document.getElementById('result');

    // 设置默认时间为当前时间
    const now = new Date();
    now.setMinutes(now.getMinutes() - now.getTimezoneOffset());
    document.getElementById('sendTime').value = now.toISOString().slice(0, 16);

    // 格式化时间为 API 需要的格式 (yyyy-mm-dd-hh-mm-ss-nanoseconds)
    function formatTimeForAPI(dateStr) {
        const date = new Date(dateStr);
        const year = date.getFullYear();
        const month = String(date.getMonth() + 1).padStart(2, '0');
        const day = String(date.getDate()).padStart(2, '0');
        const hours = String(date.getHours()).padStart(2, '0');
        const minutes = String(date.getMinutes()).padStart(2, '0');
        const seconds = String(date.getSeconds()).padStart(2, '0');
        const nanoseconds = '000000000';
        return `${year}-${month}-${day}-${hours}-${minutes}-${seconds}-${nanoseconds}`;
    }

    // 显示结果
    function showResult(success, message) {
        resultDiv.style.display = 'block';
        resultDiv.className = 'result ' + (success ? 'success' : 'error');
        resultDiv.textContent = message;
    }

    // 隐藏结果
    function hideResult() {
        resultDiv.style.display = 'none';
    }

    // 设置按钮状态
    function setButtonLoading(isLoading) {
        submitBtn.disabled = isLoading;
        btnText.style.display = isLoading ? 'none' : 'inline';
        btnLoading.style.display = isLoading ? 'inline' : 'none';
    }

    // 发送邮件
    async function sendEmail(time, subject, body) {
        const apiTime = formatTimeForAPI(time);
        
        const response = await fetch('/send', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                time: apiTime,
                subject: subject,
                body: body
            })
        });

        return await response.json();
    }

    // 表单提交
    form.addEventListener('submit', async function(e) {
        e.preventDefault();
        
        hideResult();
        setButtonLoading(true);

        const time = document.getElementById('sendTime').value;
        const subject = document.getElementById('subject').value.trim();
        const body = document.getElementById('body').value.trim();

        if (!subject || !body) {
            showResult(false, '请填写完整的信息');
            setButtonLoading(false);
            return;
        }

        try {
            const data = await sendEmail(time, subject, body);
            
            if (data.base_msg && data.base_msg.code === 200 && data.data && data.data.is_success) {
                showResult(true, '✅ 邮件发送任务已创建成功！');
                form.reset();
                // 重新设置默认时间
                const now = new Date();
                now.setMinutes(now.getMinutes() - now.getTimezoneOffset());
                document.getElementById('sendTime').value = now.toISOString().slice(0, 16);
            } else {
                showResult(false, '❌ ' + (data.base_msg?.message || '发送失败'));
            }
        } catch (error) {
            showResult(false, '❌ 网络错误，请稍后重试');
            console.error('Error:', error);
        } finally {
            setButtonLoading(false);
        }
    });
});
