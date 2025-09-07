<script setup lang="ts">
import {computed, nextTick, onMounted, ref} from "vue";
import hljs from "highlight.js";
import "highlight.js/styles/atom-one-dark.css";

const api = import.meta.env.DEV ? "http://localhost:8787" : "";

// Paste Controls
const paste_text = ref<string>("");
const currentPasteId = ref<string>("");
const currentLang = ref("");

// Hover Hints
const isHovered = ref(false);
const hintTitle = ref("")
const hintText = ref("")

// Extract the Paste ID and language (if specified) from the URL
const parseID = () => {
  const path = window.location.pathname.slice(1);
  if (!path) return [null, null];

  const parts = path.split(".");
  const id = parts[0];
  const lang = parts[1] || "";
  return [id, lang];
};

// Add a new paste
const addPaste = async () => {
  try {
    const res = await fetch(`${api}/documents`, {
      method: "POST",
      body: paste_text.value,
      headers: { "Content-Type": "text/plain" },
    });

    // API Result
    const json = await res.json();
    const key = json.key;

    // Key was found
    if (key) {
      const langSuffix = currentLang.value ? `.${currentLang.value}` : "";
      currentPasteId.value = key;

      const detected = hljs.highlightAuto(paste_text.value);
      currentLang.value = detected.language || "";

      // Update the URL
      if (detected.language) {
        window.history.pushState({}, "", `/${key}.${detected.language}`);
      } else {
        window.history.pushState({}, "", `/${key}${langSuffix}`);
      }

      // Activate code highlight
      await nextTick(() => {
        const el = document.getElementById("paste");
        if (el) hljs.highlightElement(el);
      });
    } else {
      alert(json)
    }
  } catch (err) {
    alert(err);
  }
};

// Try to load an existing text from the server
const loadPaste = async (id: string, lang: string | null) => {
  try {
    const res = await fetch(`${api}/documents/${id}`);
    if (res.ok) {
      const json = await res.json();
      paste_text.value = json.data;
      currentPasteId.value = json.key;

      // Lang code was specified
      if (lang) {
        currentLang.value = lang;
      } else {
        // Try to auto-detect language
        const detected = hljs.highlightAuto(json.data);
        currentLang.value = detected.language || "";

        // Update the URL
        if (detected.language) {
          window.history.pushState({}, "", `/${id}.${detected.language}`);
        }
      }

      // Activate code highlight
      await nextTick(() => {
        const el = document.getElementById("paste");
        if (el) hljs.highlightElement(el);
      });
    } else {
      paste_text.value = "Paste not found";
    }
  } catch (err) {
    paste_text.value = "Paste not found";
  }
};

// Handler to redirect to raw paste
const viewRaw = () => {
  window.location.href = `${api}/raw/${currentPasteId.value}`;
}

// Duplicate and edit handler
const duplicateAndEdit = () => {
  currentPasteId.value = '';
  window.history.pushState({}, "", `/`);
}

// New Paste handler
const newPaste = () => {
  currentPasteId.value = "";
  paste_text.value = "";
  window.history.pushState({}, "", `/`);
}

// On mounted hook to check for paste key and lang
onMounted(() => {
  const [id, lang] = parseID();
  if (id) loadPaste(id, lang);
});

// Line Numbers helper
const lineCount = computed(() => {
  if (!paste_text.value || currentPasteId.value == "") return 0;
  return paste_text.value.split("\n").length;
});

// Button Hint helper
const showHint = (key: string, show: boolean) => {
  isHovered.value = show;

  switch (key) {
    case "save":
      hintTitle.value = "Save";
      hintText.value = "Control or Command + s";
      break;

    case "new":
      hintTitle.value = "New";
      hintText.value = "Control or Command + n";
      break;

    case "duplicate":
      hintTitle.value = "Duplicate & Edit";
      hintText.value = "Control or Command + d";
      break;

    case "raw":
      hintTitle.value = "Just Text";
      hintText.value = "Control or Command + shift + r";
      break;
  }

}
</script>

<template>
  <div class="h-full w-full relative">
    <!-- Top Header Box -->
    <div class="header fixed top-0 right-0">
      <div class="logo-box flex items-center gap-2 justify-center">
        <img alt="icon" style="width: 20px;" src="/paste.svg" />
        <h1 class="font-bold text-lg text-white">PasteBox</h1>
      </div>

      <!-- Buttons -->
      <div class="btn-box">
        <div class="flex gap-1 items-center justify-center">
          <!-- Save Button -->
          <a href="javascript:void(0);" class="btnn" :class="paste_text.length == 0 || currentPasteId != '' ? 'pointer-events-none opacity-50 cursor-not-allowed' : ''" @click="addPaste"  @mouseenter="showHint('save', true)" @mouseleave="showHint('save', false)">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 640 640"><!--!Font Awesome Free v7.0.1 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license/free Copyright 2025 Fonticons, Inc.--><path d="M160 96C124.7 96 96 124.7 96 160L96 480C96 515.3 124.7 544 160 544L480 544C515.3 544 544 515.3 544 480L544 237.3C544 220.3 537.3 204 525.3 192L448 114.7C436 102.7 419.7 96 402.7 96L160 96zM192 192C192 174.3 206.3 160 224 160L384 160C401.7 160 416 174.3 416 192L416 256C416 273.7 401.7 288 384 288L224 288C206.3 288 192 273.7 192 256L192 192zM320 352C355.3 352 384 380.7 384 416C384 451.3 355.3 480 320 480C284.7 480 256 451.3 256 416C256 380.7 284.7 352 320 352z"/></svg>
          </a>

          <!-- New Paste Button -->
          <a href="javascript:void(0);" @click="newPaste" class="btnn" @mouseenter="showHint('new', true)" @mouseleave="showHint('new', false)">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 640 640"><!--!Font Awesome Free v7.0.1 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license/free Copyright 2025 Fonticons, Inc.--><path d="M128 64C92.7 64 64 92.7 64 128L64 512C64 547.3 92.7 576 128 576L308 576C285.3 544.5 272 505.8 272 464C272 363.4 349.4 280.8 448 272.7L448 234.6C448 217.6 441.3 201.3 429.3 189.3L322.7 82.7C310.7 70.7 294.5 64 277.5 64L128 64zM389.5 240L296 240C282.7 240 272 229.3 272 216L272 122.5L389.5 240zM464 608C543.5 608 608 543.5 608 464C608 384.5 543.5 320 464 320C384.5 320 320 384.5 320 464C320 543.5 384.5 608 464 608zM480 400L480 448L528 448C536.8 448 544 455.2 544 464C544 472.8 536.8 480 528 480L480 480L480 528C480 536.8 472.8 544 464 544C455.2 544 448 536.8 448 528L448 480L400 480C391.2 480 384 472.8 384 464C384 455.2 391.2 448 400 448L448 448L448 400C448 391.2 455.2 384 464 384C472.8 384 480 391.2 480 400z"/></svg>
          </a>

          <!-- Edit Button -->
          <a href="javascript:void(0);" @click="duplicateAndEdit" class="btnn" :class="currentPasteId == '' ? 'pointer-events-none opacity-50 cursor-not-allowed' : ''" @mouseenter="showHint('duplicate', true)" @mouseleave="showHint('duplicate', false)">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 640 640"><!--!Font Awesome Free v7.0.1 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license/free Copyright 2025 Fonticons, Inc.--><path d="M128.1 64C92.8 64 64.1 92.7 64.1 128L64.1 512C64.1 547.3 92.8 576 128.1 576L274.3 576L285.2 521.5C289.5 499.8 300.2 479.9 315.8 464.3L448 332.1L448 234.6C448 217.6 441.3 201.3 429.3 189.3L322.8 82.7C310.8 70.7 294.5 64 277.6 64L128.1 64zM389.6 240L296.1 240C282.8 240 272.1 229.3 272.1 216L272.1 122.5L389.6 240zM332.3 530.9L320.4 590.5C320.2 591.4 320.1 592.4 320.1 593.4C320.1 601.4 326.6 608 334.7 608C335.7 608 336.6 607.9 337.6 607.7L397.2 595.8C409.6 593.3 421 587.2 429.9 578.3L548.8 459.4L468.8 379.4L349.9 498.3C341 507.2 334.9 518.6 332.4 531zM600.1 407.9C622.2 385.8 622.2 350 600.1 327.9C578 305.8 542.2 305.8 520.1 327.9L491.3 356.7L571.3 436.7L600.1 407.9z"/></svg>
          </a>

          <!-- Just Text button -->
          <a href="javascript:void(0);" class="btnn" :class="currentPasteId == '' ? 'pointer-events-none opacity-50 cursor-not-allowed' : ''" @click="viewRaw"  @mouseenter="showHint('raw', true)" @mouseleave="showHint('raw', false)">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 640 640"><!--!Font Awesome Free v7.0.1 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license/free Copyright 2025 Fonticons, Inc.--><path d="M128 128C128 92.7 156.7 64 192 64L341.5 64C358.5 64 374.8 70.7 386.8 82.7L493.3 189.3C505.3 201.3 512 217.6 512 234.6L512 512C512 547.3 483.3 576 448 576L192 576C156.7 576 128 547.3 128 512L128 128zM336 122.5L336 216C336 229.3 346.7 240 360 240L453.5 240L336 122.5zM248 320C234.7 320 224 330.7 224 344C224 357.3 234.7 368 248 368L392 368C405.3 368 416 357.3 416 344C416 330.7 405.3 320 392 320L248 320zM248 416C234.7 416 224 426.7 224 440C224 453.3 234.7 464 248 464L392 464C405.3 464 416 453.3 416 440C416 426.7 405.3 416 392 416L248 416z"/></svg>
          </a>
        </div>
      </div>

      <!-- Hint Box -->
      <div class="description-box text-white" v-if="isHovered">
        <h1 class="font-bold">{{hintTitle}}</h1>
        <p class="font-light">{{hintText}}</p>
      </div>
    </div>

    <!-- Line Numbers -->
    <div class="line-numbers">
      <p v-if="currentPasteId == ''">></p>
      <div v-for="n in lineCount" :key="n">{{ n }}<br></div>
    </div>

    <!-- Paste Area -->
    <pre v-if="currentPasteId != ''" id="codebox"><code id="paste" :class="currentLang">{{ paste_text }}</code></pre>
    <textarea v-model="paste_text" v-if="currentPasteId == ''"></textarea>
  </div>
</template>
