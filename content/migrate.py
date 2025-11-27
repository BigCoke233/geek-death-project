#!/usr/bin/env python3
import os
import yaml
import unicodedata
import re

# 需要处理的目录
TARGET_DIRS = ["./struct", "./posts", "./pour"]
VALID_EXT = [".md", ".markdown"]

def auto_slug(text):
    # 全部转小写
    text = text.lower()
    # 用连字符替换空白
    text = re.sub(r"\s+", "-", text)
    # 移除大部分符号，但保留中文
    text = "".join(
        c if unicodedata.category(c).startswith(("L", "N")) else "-"
        for c in text
    )
    # 合并重复的 '-'
    text = re.sub(r"-+", "-", text)
    # 去掉头尾 '-'
    return text.strip("-")

# --------------------------------------------------------
# 读取 YAML Front Matter
# --------------------------------------------------------
def read_front_matter(content):
    """
    返回: (yaml_dict, body)
    """
    if not content.startswith("---"):
        return None, content

    end = content.find("\n---", 3)
    if end == -1:
        return None, content

    fm_raw = content[3:end].strip()
    body = content[end + 4:]

    try:
        fm = yaml.safe_load(fm_raw) or {}
        return fm, body
    except Exception:
        return None, content


# --------------------------------------------------------
# 写回 YAML Front Matter
# --------------------------------------------------------
def dump_front_matter(fm):
    return "---\n" + yaml.safe_dump(fm, sort_keys=False, allow_unicode=True) + "---\n"


# --------------------------------------------------------
# 生成 alias
# --------------------------------------------------------

def build_alias(file_path, fm):
    rel_path = os.path.relpath(file_path, ".")
    parts = rel_path.split(os.sep)

    # 忽略 _index.md
    if parts[-1].lower() == "_index.md":
        return None

    # 移除顶级目录 struct/posts/pour
    if parts[0] in ("struct", "posts", "pour"):
        parts = parts[1:]

    # 文件名（无扩展名）
    filename = os.path.splitext(parts[-1])[0]

    # slug 优先
    name = fm.get("slug", filename)

    # ---- 新逻辑：对 slug / name 做安全化与大小写处理 ----
    # 1. 空格换成连字符
    # 2. 全部转为小写
    # 3. 保留中文，保留字母数字与连字符
    name = name.strip()
    name = name.lower()  # 全部小写
    name = name.replace(" ", "-")
    # 只去除 *不安全* 的字符，避免破坏中文
    name = re.sub(r"[^a-z0-9\-\u4e00-\u9fff]", "", name)

    # 构建路径
    if len(parts) > 1:
        sub = "/".join(parts[:-1])
        return f"/posts/{sub}/{name}"
    else:
        return f"/posts/{name}"




def should_ignore(path):
    """忽略 _index.md / _index.markdown，不区分大小写"""
    filename = os.path.basename(path).lower()
    name, ext = os.path.splitext(filename)
    if name == "_index" and ext in VALID_EXT:
        return True
    return False

# --------------------------------------------------------
# 处理单个 Markdown 文件
# --------------------------------------------------------
def process_file(path):
    if should_ignore(path):
        return

    with open(path, "r", encoding="utf-8") as f:
        content = f.read()

    fm, body = read_front_matter(content)
    if fm is None:
        print(f"[跳过] 无 YAML front-matter: {path}")
        return

    alias = build_alias(path, fm)

    # 添加 aliases 字段
    if "aliases" not in fm:
        fm["aliases"] = [alias]
    else:
        if not isinstance(fm["aliases"], list):
            fm["aliases"] = [fm["aliases"]]
        if alias not in fm["aliases"]:
            fm["aliases"].append(alias)

    # 写回文件
    new_content = dump_front_matter(fm) + body
    with open(path, "w", encoding="utf-8") as f:
        f.write(new_content)

    print(f"[更新] {path} → {alias}")


# --------------------------------------------------------
# 主函数
# --------------------------------------------------------
def main():
    for root_dir in TARGET_DIRS:
        for root, dirs, files in os.walk(root_dir):
            for file in files:
                if os.path.splitext(file)[1].lower() in VALID_EXT:
                    full = os.path.join(root, file)
                    process_file(full)


if __name__ == "__main__":
    main()
