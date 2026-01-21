// implement features of like buttons in action menu
// 1. fetch likes
// 2. send like to server
// 3. visualize like data (orgam farm)
// 4. update like states

const LIKE_SERVER_URL = "https://like.eltr.ac/"

function initLikeButton() {
    if (window.likeButtonInitialized) {
        return;
    }
    
    const likeBtn = document.getElementById("like-btn");
    if (!likeBtn) {
        return;
    }

    const pageId = likeBtn.dataset.pageid;;

    // get required elements
    const elements = {
        count: document.getElementById("like-count-" + pageId),
        btn: document.getElementById("like-btn"), // Main button doesn't have page ID suffix
        container: document.getElementById("like-container-" + pageId),
        icon: document.getElementById("like-icon-" + pageId)
    };
    // return if no page ID found or not in production or no elements
    if (
        !pageId || !elements.count || !elements.btn || !elements.container
    ) {
        return;
    }

    function setLikedState() {
        // update classes
        elements.icon?.classList.add('liked-heart');
        elements.icon?.classList.add('animate-heart-beat');
        elements.count?.classList.add('liked-count');
        elements.btn.disabled = true;
        elements.container.classList.add('liked');
        // update tooltip
        elements.btn.setAttribute('title',
          elements.container.classList.contains('liked')
            ? elements.btn.getAttribute('data-liked-title')
            : '为这篇文章献上心脏'
        );
    }

    async function fetchLikes() {
        try {
          const response = await fetch(LIKE_SERVER_URL + "like/" + pageId);
          const data = await response.json();
          elements.count.innerText = data.likes ?? 0;
          if (localStorage.getItem("liked_" + pageId)) {
            setLikedState();
            updateOrganFarm(data.likes);
          }
        } catch(e) {
          console.log(e);
          elements.count.innerText = "❌";
        }
    }

    async function sendLike(pageId) {
        if (localStorage.getItem("liked_" + pageId)) return;

        try {
          const response = await fetch(LIKE_SERVER_URL + "like/" + pageId, { method: "POST" });
          const data = await response.json();
          localStorage.setItem("liked_" + pageId, "1");
          elements.count.innerText = data.likes ?? 1;
          updateOrganFarm(data.likes);
          setLikedState();
        } catch(e) {
          console.log(e);
          alert("点赞失败，请稍后再试");
        }
    };

    function updateOrganFarm(count) {
        const organFarm = document.getElementById("organ-farm");
        const organCount = count || parseInt(elements.count.innerText);
        const organImages = [
          "https://r2.eltr.ac/picgo2025/oragn-3.png",
          "https://r2.eltr.ac/picgo2025/organ-1.png",
          "https://r2.eltr.ac/picgo2025/organ-2.png"
        ];

        if (!organFarm) return;
        if (organFarm.classList.contains("hidden")) {
          organFarm.classList.remove("hidden");
          organFarm.classList.add("animate-slide-in-down");
        }

        organFarm.innerHTML = "";
        for (let i = 0; i < organCount; i++) {
          const img = document.createElement("img");
          img.className = "transition hover:scale-150 dark:invert";

          const sizeRange = [6, 10];
          const size = Math.floor(Math.random() * (sizeRange[1] - sizeRange[0] + 1)) + sizeRange[0];
          img.classList.add(`w-${size}`, `h-${size}`);

          img.src = organImages[Math.floor(Math.random() * organImages.length)];
          // randomly offset and rotate
          img.classList.add("rotate-" + Math.floor(Math.random() * 5));
          img.classList.add("translate-x-" + Math.floor(Math.random() * 5));
          img.classList.add("translate-y-" + Math.floor(Math.random() * 5));
          img.classList.add("opacity" + (100 - Math.floor(Math.random() * 3)*10));

          organFarm.appendChild(img);
        }
    }

    // fetch likes when action menu is visible
    const actionMenuEl = document.getElementById("action-menu");
    if (actionMenuEl) {
        const observer = new IntersectionObserver((entries, obs) => {
            if (entries[0].isIntersecting) {
              fetchLikes();
              obs.disconnect();
            }
        });
        observer.observe(actionMenuEl);
    } else {
        fetchLikes();
    }

    // add click event listener to like button
    likeBtn.addEventListener("click", () => {
        sendLike(pageId)
    })
    
    window.likeButtonInitialized = true;
}

// Initialize when DOM is ready
if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', initLikeButton);
} else {
    initLikeButton();
}

// Fallback initialization
window.addEventListener('load', () => {
    if (!window.likeButtonInitialized) {
        initLikeButton();
    }
});
