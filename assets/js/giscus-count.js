// count giscus comments and update DOM
// author: Eltrac
document.addEventListener('DOMContentLoaded', () => {
    const commentCountEl = document.getElementById('comment-count');
    if (!commentCountEl) return;

    const observer = new MutationObserver(() => {
      // check if giscus is loaded
      const giscusIframe = document.querySelector('iframe.giscus-frame');
      if (!giscusIframe) return;

      giscusIframe.addEventListener('load', () => {
        // receive message from giscus iframe
        window.addEventListener('message', function(event) {
          if (event.origin !== 'https://giscus.app') return;
          const giscusData = event.data;
          if (giscusData.giscus?.discussion) {
            const discussion = giscusData.giscus.discussion;
            const totalComments = (discussion.totalCommentCount || 0) + (discussion.totalReplyCount || 0);
            // set comment count
            commentCountEl.textContent = totalComments.toString();
          // if no discussion found, set to 0
          } else commentCountEl.textContent = '0';
        });
      });

      observer.disconnect();
    });
    observer.observe(document.body, { childList: true, subtree: true });
});
