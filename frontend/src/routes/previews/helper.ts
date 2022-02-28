import "prismjs/components/prism-bash";
import "prismjs/components/prism-batch";
import "prismjs/components/prism-c";
import "prismjs/components/prism-c";
import "prismjs/components/prism-clojure";
import "prismjs/components/prism-csv";
import "prismjs/components/prism-d";
import "prismjs/components/prism-diff";
import "prismjs/components/prism-docker";
import "prismjs/components/prism-elm";
import "prismjs/components/prism-gcode";
import "prismjs/components/prism-go";
import "prismjs/components/prism-haskell";
import "prismjs/components/prism-ignore";
import "prismjs/components/prism-ini";
import "prismjs/components/prism-java";
import "prismjs/components/prism-javascript";
import "prismjs/components/prism-json";
import "prismjs/components/prism-jsx";
import "prismjs/components/prism-julia";
import "prismjs/components/prism-kotlin";
import "prismjs/components/prism-latex";
import "prismjs/components/prism-lisp";
import "prismjs/components/prism-list";
import "prismjs/components/prism-lua";
import "prismjs/components/prism-markup-templating"
import "prismjs/components/prism-nix";
import "prismjs/components/prism-php";
import "prismjs/components/prism-powershell";
import "prismjs/components/prism-python";
import "prismjs/components/prism-qml";
import "prismjs/components/prism-r";
import "prismjs/components/prism-ruby";
import "prismjs/components/prism-rust";
import "prismjs/components/prism-sass";
import "prismjs/components/prism-scss";
import "prismjs/components/prism-sql";
import "prismjs/components/prism-swift";
import "prismjs/components/prism-tsx";
import "prismjs/components/prism-typescript";
import "prismjs/components/prism-vala";
import "prismjs/components/prism-wasm";
import "prismjs/components/prism-yaml";

const extensions = new Map<string, string>([
    ["bat", "batch"],
    ["c", "c"],
    ["cl", "lisp"],
    ["clj", "clojure"],
    ["csv", "csv"],
    ["d", "d"],
    ["diff", "diff"],
    ["dockerfile", "docker"],
    ["el", "emacs"],
    ["elm", "elm"],
    ["emacs", "emacs"],
    ["gcode", "gcode"],
    ["gitignore", "ignore"],
    ["go", "go"],
    ["h", "c"],
    ["hs", "haskell"],
    ["ini", "ini"],
    ["java", "java"],
    ["jl", "julia"],
    ["js", "javascript"],
    ["json", "json"],
    ["jsx", "jsx"],
    ["kt", "kotlin"],
    ["l", "lisp"],
    ["lisp", "lisp"],
    ["lua", "lua"],
    ["nix", "nix"],
    ["php", "php"],
    ["ps1", "powershell"],
    ["psm1", "powershell"],
    ["py", "python"],
    ["qml", "qml"],
    ["r", "r"],
    ["rb", "ruby"],
    ["rs", "rust"],
    ["sass", "sass"],
    ["scala", "scala"],
    ["scss", "scss"],
    ["sh", "shell"],
    ["sql", "sql"],
    ["swift", "swift"],
    ["tex", "latex"],
    ["ts", "typescript"],
    ["tsx", "tsx"],
    ["vala", "vala"],
    ["wasm", "wasm"],
    ["yaml", "yaml"],
    ["yml", "yaml"],
]);

export function getLanguage(name: string): string {
    let ext = name.split('.').pop();

    if (extensions.has(ext)) {
        return extensions.get(ext);
    } else {
        return "plain";
    }
}