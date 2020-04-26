# x86-64

[X86-64(en.wikipedia)](https://en.wikipedia.org/wiki/X86-64)

x86-64 (x64、x86_64、AMD64、Intel 64 としても知られています) は、x86 命令セットの 64 ビット版です。64ビットモードと互換性モードの2つの新しい動作モードと、新しい4レベルのページングモードが導入されています。64ビットモードと新しいページングモードでは、32ビット版の前身と比べて大容量の仮想メモリと物理メモリをサポートし、プログラムがより大容量のデータをメモリに格納できるようになっています。浮動小数点演算は必須のSSE2ライクな命令でサポートされており、x87/MMXスタイルのレジスタは一般的には使用されません（64ビットモードでも使用可能です）。その代わりに、32個のベクトルレジスタのセットが使用されます(各レジスタには、1つまたは2つの倍精度の数値、または1～4つの単精度の数値、またはさまざまな整数形式を格納できます)。64ビットモードでは、命令は64ビットオペランドと64ビットアドレッシングモードをサポートするように変更されます。互換モードでは、64ビット・オペレーティング・システムがサポートしている場合、16ビットおよび32ビットのユーザー・アプリケーションを64ビット・アプリケーションと変更せずに共存させて実行することができます 、完全なx86 16ビットおよび32ビット命令セットがハードウェアに実装されたままであるため、これらの古い実行可能ファイルは性能の低下をほとんど感じずに実行できます。また、x86-64をサポートするプロセッサは、80286以降のx86プロセッサと同様に、完全な下位互換性を実現するためにリアルモードで電源を入れます。

AMDが作成し、2000年にリリースされたオリジナル仕様は、AMD、Intel、VIAの3社が実装している。OpteronやAthlon 64プロセッサに搭載されたAMDのK8マイクロアーキテクチャが最初に実装された。これは、Intel以外の会社が設計したx86アーキテクチャへの最初の大きな追加だった。Intelは追随を余儀なくされ、AMDの仕様とソフトウェア互換性のある修正NetBurstファミリを発表した。VIA Technologiesは、VIA IsaiahアーキテクチャのVIA Nanoでx86-64を導入した。

x86-64アーキテクチャは、Intel Itaniumアーキテクチャ(旧IA-64)とは一線を画しており、x86アーキテクチャとはネイティブ命令セットレベルでの互換性がない。一方のアーキテクチャ用にコンパイルされたオペレーティングシステムやアプリケーションは、もう一方のアーキテクチャでは実行できません。



# POSIX

[POSIX(en.wikipedia)](https://en.wikipedia.org/wiki/POSIX)

POSIX（Portable Operating System Interface）は、オペレーティングシステム間の互換性を維持するためにIEEE Computer Societyによって規定された規格のファミリーです。 POSIXは、Unixや他のオペレーティングシステムの亜種とのソフトウェア互換性のために、コマンドラインシェルやユーティリティインターフェースとともに、アプリケーションプログラミングインターフェース（API）を定義します。



### 名前
もともと「POSIX」という名称は、1988年に発表されたIEEE Std 1003.1-1988を指していました。POSIX規格のファミリーは、正式にはIEEE 1003として指定され、国際規格名はISO/IEC 9945です。

規格は1985年頃に始まったプロジェクトから生まれました。リチャード・ストールマンは、以前のIEEE-IXではなく、POSIXという名前をIEEEに提案しました。委員会は、その方が発音しやすく記憶に残りやすいと判断し、これを採用した。



### 概要
Unixが標準システムインターフェースの基礎として選ばれたのは、「メーカーニュートラル」という理由もありました。しかし、Unixにはいくつかの主要なバージョンが存在していたため、共通分母システムを開発する必要がありました。UnixライクなオペレーティングシステムのためのPOSIX仕様は、もともとはコアプログラミングインターフェイスのための単一の文書で構成されていたが、最終的には19の別個の文書(POSIX.1、POSIX.2など)にまで成長した。POSIXはまた、ほとんどの最新のオペレーティングシステムでサポートされている標準のスレッディングライブラリAPIを定義しています。2008年、POSIXのほとんどの部分が単一の規格にまとめられました (IEEE Std 1003.1-2008、POSIX.1-2008としても知られています)。

2014年現在、POSIXのドキュメントは2つの部分に分かれています。

POSIX.1、2013年版。POSIX基本定義、システムインタフェース、コマンドとユーティリティ（POSIX.1、POSIX.1の拡張機能、リアルタイムサービス、スレッドインタフェース、リアルタイム拡張機能、セキュリティインタフェース、ネットワークファイルアクセスとネットワークプロセス間通信、ユーザーポータビリティ拡張機能、修正と拡張機能、保護と制御ユーティリティ、バッチシステムユーティリティを含む。これは、POSIX 1003.1-2008 技術解説書1付きのものです)。
POSIX適合性試験。POSIX適合性試験: POSIXのための試験スイートは標準に付属している。VSX-PCTS または VSX POSIX 適合性テストスイート。
POSIX規格の開発は、オースティングループ（IEEE、The Open Group、ISO/IEC JTC 1の共同作業グループ）で行われています。



# ユーザランド

[User_space#USERLAND(en.wikipedia)](https://en.wikipedia.org/wiki/User_space#USERLAND)

ユーザランド(またはユーザ空間)という用語は、オペレーティングシステムのカーネルの外で動作するすべてのコードを指します。

各ユーザ空間プロセスは通常、それ自身の仮想メモリ空間で動作し、明示的に許可されていない限り、他のプロセスのメモリにアクセスすることはできません。これは、今日の主流のオペレーティングシステムにおけるメモリ保護の基礎であり、特権分離のための構成要素でもあります。分離されたユーザモードは、効率的な仮想マシンを構築するために使用することもできます - Popek and Goldberg の仮想化要件を参照してください。十分な権限があれば、プロセスはデバッガの場合と同様に、他のプロセスのメモリ空間の一部を自分のメモリ空間にマップするようカーネルに要求することができます。プログラムは他のプロセスとの共有メモリ領域を要求することもできますが、プロセス間の通信を可能にする他の技術も利用できます。



# X11



# Wayland



# GNOME



# KDE Plasma



# x86

https://en.wikipedia.org/wiki/X86#x86_registers

x86は、Intel 8086マイクロプロセッサとその8088変種をベースにIntelによって最初に開発された命令セット・アーキテクチャのファミリ[a]である。8086は、インテルの8ビット・マイクロプロセッサ8080の完全な16ビット拡張として1978年に導入され、プレーンな16ビット・アドレスでカバーできるより多くのメモリをアドレス指定するためのソリューションとしてメモリ・セグメンテーションが採用された。x86」という用語が生まれたのは、インテルの8086プロセッサの後継となる80186、80286、80386、80486プロセッサを含むいくつかのプロセッサの名前が「86」で終わることに由来する。

このアーキテクチャは、Intel、Cyrix、AMD、VIA Technologies、その他多くの企業のプロセッサに実装されており、Zet SoC プラットフォーム (現在は非アクティブ) のようなオープンな実装もあります[2]。 しかしながら、これらのうち、x86 アーキテクチャのライセンスを保持しているのは Intel、AMD、VIA Technologies、DM&P Electronics だけであり、これらの中から最新の 64 ビット設計を積極的に生産しているのは、最初の 2 社だけです。

この用語はIBM PC互換性と同義ではないが、これは他の多数のコンピュータ・ハードウェアを意味しているためである; 組み込みシステムや汎用コンピュータは、PC互換市場が始まる前からx86チップを使用しており[c]、そのうちのいくつかはIBM PC(1981年)そのものよりも前から使用されていたものである。

2018年現在、販売されているパーソナルコンピュータやラップトップの大半はx86アーキテクチャをベースにしているが、他のカテゴリー、特にスマートフォンやタブレットなどの大量[解明が必要]なモバイルカテゴリーはARMが支配しており、ハイエンドでは、x86は引き続き計算集約型ワークステーションやクラウドコンピューティングのセグメントを支配している[3]。



| Designer                                                     | [Intel](https://en.wikipedia.org/wiki/Intel), [AMD](https://en.wikipedia.org/wiki/Advanced_Micro_Devices) |
| :----------------------------------------------------------- | ------------------------------------------------------------ |
| Bits                                                         | [16-bit](https://en.wikipedia.org/wiki/16-bit), [32-bit](https://en.wikipedia.org/wiki/32-bit) and [64-bit](https://en.wikipedia.org/wiki/64-bit_computing) |
| Introduced                                                   | 1978 (16-bit), 1985 (32-bit), 2003 (64-bit)                  |
| [Design](https://en.wikipedia.org/wiki/Computer_architecture) | [CISC](https://en.wikipedia.org/wiki/Complex_instruction_set_computing) |
| Type                                                         | [Register–memory](https://en.wikipedia.org/wiki/Register_memory_architecture) |
| [Encoding](https://en.wikipedia.org/wiki/Instruction_set)    | Variable (1 to 15 bytes)                                     |
| [Branching](https://en.wikipedia.org/wiki/Branch_(computer_science)) | [Condition code](https://en.wikipedia.org/wiki/Status_register) |
| [Endianness](https://en.wikipedia.org/wiki/Endianness)       | Little                                                       |
| Page size                                                    | [8086](https://en.wikipedia.org/wiki/8086)–[i286](https://en.wikipedia.org/wiki/I286): None [i386](https://en.wikipedia.org/wiki/Intel_386), [i486](https://en.wikipedia.org/wiki/Intel_486): 4 KB pages [P5](https://en.wikipedia.org/wiki/P5_(microarchitecture)) [Pentium](https://en.wikipedia.org/wiki/Pentium_(brand)): added 4 MB pages (Legacy [PAE](https://en.wikipedia.org/wiki/Physical_Address_Extension): 4 KB→2 MB) [x86-64](https://en.wikipedia.org/wiki/Long_mode): added 1 GB pages |
| Extensions                                                   | [x87](https://en.wikipedia.org/wiki/X87), [IA-32](https://en.wikipedia.org/wiki/IA-32), [x86-64](https://en.wikipedia.org/wiki/X86-64), [MMX](https://en.wikipedia.org/wiki/MMX_(instruction_set)), [3DNow!](https://en.wikipedia.org/wiki/3DNow!), [SSE](https://en.wikipedia.org/wiki/Streaming_SIMD_Extensions), [SSE2](https://en.wikipedia.org/wiki/SSE2), [SSE3](https://en.wikipedia.org/wiki/SSE3), [SSSE3](https://en.wikipedia.org/wiki/SSSE3), [SSE4](https://en.wikipedia.org/wiki/SSE4), [SSE4.2](https://en.wikipedia.org/wiki/SSE4.2), [SSE5](https://en.wikipedia.org/wiki/SSE5), [AES-NI](https://en.wikipedia.org/wiki/AES-NI), [CLMUL](https://en.wikipedia.org/wiki/CLMUL_instruction_set), [RDRAND](https://en.wikipedia.org/wiki/RDRAND), [SHA](https://en.wikipedia.org/wiki/Intel_SHA_extensions), [MPX](https://en.wikipedia.org/wiki/Intel_MPX), [SGX](https://en.wikipedia.org/wiki/Software_Guard_Extensions), [XOP](https://en.wikipedia.org/wiki/XOP_instruction_set), [F16C](https://en.wikipedia.org/wiki/F16C), [ADX](https://en.wikipedia.org/wiki/Intel_ADX), [BMI](https://en.wikipedia.org/wiki/Bit_Manipulation_Instruction_Sets), [FMA](https://en.wikipedia.org/wiki/FMA_instruction_set), [AVX](https://en.wikipedia.org/wiki/Advanced_Vector_Extensions), [AVX2](https://en.wikipedia.org/wiki/AVX2), [AVX512](https://en.wikipedia.org/wiki/AVX512), [VT-x](https://en.wikipedia.org/wiki/VT-x), [AMD-V](https://en.wikipedia.org/wiki/AMD-V), [TSX](https://en.wikipedia.org/wiki/Transactional_Synchronization_Extensions), [ASF](https://en.wikipedia.org/wiki/Advanced_Synchronization_Facility) |
| Open                                                         | Partly. For some advanced features, x86 may require license from Intel; x86-64 may require an additional license from AMD. The 80486 processor has been on the market for more than 20 years[[1\]](https://en.wikipedia.org/wiki/X86#cite_note-1) and so cannot be subject to patent claims. The pre-586 subset of the x86 architecture is therefore fully open. |
| [Registers](https://en.wikipedia.org/wiki/Processor_register) |                                                              |
| [General purpose](https://en.wikipedia.org/wiki/General_purpose_register) | 16-bit: 6 semi-dedicated registers, BP and SP are not general-purpose32-bit: 8 GPRs, including EBP and ESP64-bit: 16 GPRs, including RBP and RSP |
| [Floating point](https://en.wikipedia.org/wiki/Floating_point) | 16-bit: optional separate [x87](https://en.wikipedia.org/wiki/X87) FPU32-bit: optional separate or integrated [x87](https://en.wikipedia.org/wiki/X87) FPU, integrated [SSE2](https://en.wikipedia.org/wiki/SSE2) units in later processors64-bit: integrated [x87](https://en.wikipedia.org/wiki/X87) and [SSE2](https://en.wikipedia.org/wiki/SSE2) units, later implementations extended to [AVX2](https://en.wikipedia.org/wiki/AVX2) and [AVX512](https://en.wikipedia.org/wiki/AVX512) |

# x86 assembly language

https://en.wikipedia.org/wiki/X86_assembly_language

x86アセンブリ言語は、1972年4月に導入されたIntel 8008にさかのぼってある程度の互換性を提供する下位互換性のあるアセンブリ言語のファミリーです。他のアセンブリ言語と同様に、短いニーモニックを使用して、コンピュータのCPUが理解して従うことができる基本的な命令を表現します。コンパイラは、高レベルのプログラムをマシンコードに変換する際に、中間的なステップとしてアセンブリコードを生成することがあります。プログラミング言語として考えられているアセンブリコードは、マシン固有の低レベルなものです。アセンブリ言語は、小規模なリアルタイム組み込みシステムやオペレーティングシステムのカーネルやデバイスドライバなど、詳細で時間的に重要なアプリケーションに使用されます。



## ニーモニックとオペコード

各 x86 アセンブリ命令はニーモニックで表現されており、これは多くの場合、1 つ以上のオペランドと組み合わされて、オペコードと呼ばれる 1 つ以上のバイトに変換されます。文書化されたニーモニックがない潜在的なオペコードもあり、これらを使用したプログラムはプロセッサによって解釈が異なるため、プログラムの動作に矛盾が生じたり、プロセッサによっては例外が発生したりすることがあります。これらのオペコードは、コードを小さくしたり、速くしたり、エレガントにしたり、作者の腕前を誇示したりするための方法として、コードを書くコンテストによく登場します。



## 文法

x86アセンブリ言語には、2つの主要な構文分岐があります。Intel 構文は元々 x86 プラットフォームの文書化に使用されていたもので、AT&T 構文は AT&T で作成されたものです。Intel構文はDOSとWindowsの世界では支配的であり、UnixはAT&T Bell Labsで作られたので、AT&T構文はUnixの世界では支配的です。ここでは、Intel 構文と AT&T 構文の主な違いをまとめてみました。

|                             AT&T                             |                            Intel                             |                                                              |
| :----------------------------------------------------------: | :----------------------------------------------------------: | ------------------------------------------------------------ |
|                       Parameter order                        |       Source before the destination.  `mov $5, %eax `        | Destination before source. `mov eax, 5 `                     |
|                        Parameter size                        | Mnemonics are suffixed with a letter indicating the size of the operands: *q* for qword, *l* for long (dword), *w* for word, and *b* for byte.[[1\] z](https://en.wikipedia.org/wiki/X86_assembly_language#cite_note-GASvsNASM-1)  `addl $4, %esp ` | Derived from the name of the register that is used (e.g. *rax, eax, ax, al* imply *q, l, w, b*, respectively). `add esp, 4 ` |
| [Sigils](https://en.wikipedia.org/wiki/Sigil_(computer_programming)) | [Immediate values](https://en.wikipedia.org/wiki/Constant_(programming)) prefixed with a "$", registers prefixed with a "%".[[1\]](https://en.wikipedia.org/wiki/X86_assembly_language#cite_note-GASvsNASM-1) | The assembler automatically detects the type of symbols; i.e., whether they are registers, constants or something else. |
| Effective [addresses](https://en.wikipedia.org/wiki/Memory_address) | General syntax of *DISP(BASE,INDEX,SCALE)*. Example:`movl mem_location(%ebx,%ecx,4), %eax ` | Arithmetic expressions in square brackets; additionally, size keywords like *byte*, *word*, or *dword* have to be used if the size cannot be determined from the operands.[[1\]](https://en.wikipedia.org/wiki/X86_assembly_language#cite_note-GASvsNASM-1) Example:`mov eax, [ebx + ecx*4 + mem_location]` |

## Segmented addressing

実質および仮想 8086 モードの x86 アーキテクチャは、他の多くの環境で使用されているフラット・メモリ・モデルではなく、セグメンテーションとして知られているプロセスを使用してメモリ・アドレスを指定します。セグメンテーションでは、セグメントとオフセットの2つの部分からメモリ・アドレスを構成します。セグメントは64KBのアドレス・グループの先頭を指し、オフセットはこの先頭アドレスから目的のアドレスまでの距離を決定します。セグメント化されたアドレスでは、完全なメモリアドレスを得るために2つのレジスタが必要になります。1つはセグメントを保持し、もう1つはオフセットを保持します。フラットアドレスに変換するためには、セグメント値を左に4ビットシフトして(24または16の乗算に相当)、オフセットに加算して完全なアドレスを形成します。

例えば、リアルモード/プロテクトのみの場合、DSに16進数の0xDEADとDXに0xCAFEが含まれているとすると、メモリアドレス0xDEAD * 0x10 + 0xCAFE = 0xEB5CEを指すことになります。したがって、CPUは実モードで最大1,048,576バイト(1MB)までアドレスを指定することができます。セグメント値とオフセット値を組み合わせることで、20ビットのアドレスが得られます。

オリジナルのIBM PCでは、プログラムは640KBに制限されていましたが、拡張メモリ仕様はバンクスイッチング方式を実装するために使用されていましたが、Windowsなどの後のオペレーティングシステムが新しいプロセッサのより大きなアドレス範囲を使用し、独自の仮想メモリ方式を実装したときには使用されなくなりました。

プロテクトモードは、Intel 80286で始まったもので、OS/2で利用されていました。BIOSにアクセスできなかったり、プロセッサをリセットしないとリアルモードに切り替えられなかったりと、いくつかの欠点がありましたが、普及するには至りませんでした[6]。80286の拡張機能にアクセスするには、オペレーティング・システムがプロセッサをプロテクト・モードに設定し、24ビット・アドレッシングを可能にして、224バイト（16メガバイト）のメモリにアクセスする必要がありました。

プロテクトモードでは、セグメントセレクタは13ビットのインデックス、エントリがGDTかLDTかを決定するテーブル・インジケータ・ビット、2ビットの要求特権レベルの3つの部分に分割されます。

セグメントとオフセットを持つアドレスを参照する場合は、セグメント:オフセットの表記が使用されます。上記の例では、フラットアドレス0xEB5CEは0xDEAD:0xCAFEとして、またはセグメントとオフセットのレジスタペア(DS:DX)として記述できます。

重要なアドレスを指すセグメントレジスタと一般レジスタの特殊な組み合わせがあります。

CS:IP (CSはコードセグメント、IPは命令ポインタ)は、プロセッサが次のバイトのコードをフェッチするアドレスを指します。
SS:SP (SS は Stack Segment, SP は Stack Pointer) は、スタックの先頭のアドレス、つまり最近プッシュされたバイトを指します。
DS:SI (DS is Data Segment, SI is Source Index) は、ES:DI にコピーされようとしている文字列データを指すのによく使われます。
ES:DI (ES はエクストラセグメント、DI はデスティネーションインデックス) は、上述したように、通常、文字列コピーの宛先を指すために使用されます。
Intel 80386には、リアルモード、プロテクトモード、バーチャルモードの3つの動作モードが搭載されていた。80286でデビューしたプロテクト・モードは、最大4GBのメモリ・アドレス指定を可能にするために拡張され、新しいバーチャル8086モード(VM86)は、1つ以上のリアル・モード・プログラムをプロテクト環境で実行できるようにした。

80386の拡張保護モードの32ビット・フラット・メモリ・モデルは、AMDが2003年にx86-64をリリースするまでのx86プロセッサ・ファミリーにとって最も重要な機能変更であったかもしれません。

# Mnemonic

https://en.wikipedia.org/wiki/Mnemonic

ニーモニック（/nəmˈmɒnɪk/、[1]最初の "m "は発音されない）デバイス、またはメモリデバイスは、人間の記憶の中で情報の保持または検索（記憶）を補助する任意の学習技術である。記憶装置は、効率的な保存と検索を可能にする方法で任意の情報を符号化するための特定のツールとして、精巧な符号化、検索キュー、およびイメージを使用しています。ニーモニックは、元の情報が、よりアクセスしやすく、意味のあるものと関連づけられるように支援します。

一般的に使用されるニーモニクスは、リストや短い詩、頭文字、頭文字、または記憶に残るフレーズなどの聴覚形式で使用されることが多いですが、ニーモニクスは他のタイプの情報や視覚的または感覚的な形式でも使用することができます。ニーモニックの使用は、人間の心がより簡単に空間的、個人的、意外性、物理的、性的、ユーモア、またはそうでなければ「関連性のある」情報を覚えているという観察に基づいています。

ニーモニック」という言葉は、古代ギリシャ語のμνημονικός (mnēmonikos)に由来し、「記憶の、記憶に関連する」[2] を意味し、ギリシャ神話に登場する記憶の女神の名前であるMnemosyne ("remembrance")と関連しています。これらの単語はいずれもμνήμη（mnēmē）、「記憶、記憶」に由来している[3] 古代における記憶術は、今日の記憶術として知られているものの文脈で考えられることがほとんどであった。

古代ギリシャ人とローマ人は、記憶には「自然な」記憶と「人工的な」記憶の2つのタイプがあると区別していた。前者は先天的なもので、誰もが本能的に使うものです。後者は対照的に訓練され、さまざまなニーモニック技術の学習と練習によって開発されなければなりません。

ニーモニックシステムは、記憶力を向上させるために意識的に使用される技術や戦略です。それらは、すでに長期記憶に保存されている情報を使用して、暗記をより簡単な作業にするのに役立つ[4]。

![img](https://upload.wikimedia.org/wikipedia/commons/thumb/9/9b/Month_-_Knuckles_%28en%29.svg/2880px-Month_-_Knuckles_%28en%29.svg.png)



# Opcode

https://en.wikipedia.org/wiki/Opcode

コンピューティングでは、オペコード（オペレーションコードの略で、命令マシンコード、命令コード、命令シラブル、命令パーセル、またはオペストリングとしても知られています）は、実行される操作を指定する機械語命令の部分です。オペコード自体の他に、ほとんどの命令はオペランドの形で処理するデータを指定します。ハードウェアデバイスである各種CPUの命令セットアーキテクチャで使用されているオペコードのほか、抽象的な計算機ではバイトコード仕様の一部として使用されることもある。

## 概要

オプコードの仕様とフォーマットは、問題のプロセッサの命令セット・アーキテクチャ（ISA）に記載されています（一般的なCPUである場合もあれば、より特殊な処理ユニットである場合もあります）。ある命令セットのオペコードは、可能なすべてのオペコードバイトを詳細に記述したオペコードテーブルを使用して記述することができます。オペコード自体とは別に、命令は通常、操作が動作すべきオペランド(すなわちデータ)のための1つ以上の指定子を持っていますが、操作によっては暗黙のオペランドを持っていたり、全く持っていなかったりすることもあります。命令セットには、オペコードとオペランド指定子のためのほぼ均一なフィールドを持つものと、より複雑で可変長の構造を持つもの(例えばx86アーキテクチャ)があります。命令セットは、予約されたバイト列に続く既存のオペコードで構成された新しい命令のサブセットを追加するオペコードプレフィックスを使用して拡張することができます。

### オペランド
アーキテクチャによっては、オペランドはレジスタ値、スタック内の値、他のメモリ値、I/Oポート（メモリマップされている場合もある）などであり、多かれ少なかれ複雑なアドレッシングモードを使用して指定し、アクセスすることができる[citation needed]。

アセンブリ言語、または単なるアセンブリは、ニーモニック、命令、オペランドを用いて機械コードを表現する低レベルのプログラミング言語である。これにより、機械命令を正確に制御しながら可読性を高めることができます。現在、ほとんどのプログラミングは高レベルのプログラミング言語を使って行われていますが、これは一般的に読み書きが簡単です。これらの言語は、システム固有のコンパイラでコンパイル（アセンブリ言語に翻訳）されるか、他のコンパイル済みプログラムを通して実行される必要があります。



# AMD(Company)

https://en.wikipedia.org/wiki/Advanced_Micro_Devices

アドバンスト・マイクロ・デバイセズ（AMD）は、カリフォルニア州サンタクララ郡に本社を置くアメリカの多国籍半導体企業です。(AMD) は、カリフォルニア州サンタクララ郡に本社を置くアメリカの多国籍半導体企業で、ビジネスおよびコンシューマ市場向けのコンピュータ・プロセッサおよび関連技術を開発しています。当初は自社でプロセッサを製造していましたが、2009年にGlobalFoundriesが分社化された後は、ファブレス化として知られる製造を外部に委託しています。AMDの主な製品には、サーバー、ワークステーション、パソコン、組み込みシステム・アプリケーション向けのマイクロプロセッサ、マザーボード・チップセット、組み込みプロセッサ、グラフィックス・プロセッサなどがあります。



# x64 vs ARM64

https://docs.microsoft.com/answers/questions/10614/what-is-the-difference-between-x64-and-arm64.html

x86やx64と同様に、ARMは別のプロセッサ(CPU)アーキテクチャです。ARMアーキテクチャは通常、モバイルデバイス用のCPUを構築するために使用されますが、ARM64は単に64ビット処理をサポートするARMアーキテクチャの拡張または進化です。ARM64アーキテクチャで構築されたデバイスには、デスクトップPC、モバイルデバイス、および一部のIoTコアデバイス（Rasperry Pi 2、Raspberry Pi 3、DragonBoard）が含まれます。例えば、Microsoft HoloLens 2はARM64プロセッサを使用しています。



# RISC

https://en.wikipedia.org/wiki/Reduced_instruction_set_computer

縮小命令セットコンピュータ、またはRISC（/rɪsk/）とは、コンピュータのマイクロプロセッサが複合命令セットコンピュータ（CISC）よりも命令あたりのサイクル数（CPI）を少なくできるようにしたコンピュータ命令セットのことである[1]。

RISCコンピュータは、複雑で専門的な命令の大規模なセットよりも、単純で一般的な命令の小さなセットを持っています。RISCの主な特徴は、命令セットが非常に規則的な命令パイプラインの流れに最適化されていることである[2]。 もう一つの共通のRISCの特徴は、メモリがほとんどの命令の一部としてではなく、特定の命令を通してアクセスされるロード/ストア・アーキテクチャ[3]である。

1960年代と1970年代の多くのコンピュータがRISCの前身であることが確認されているが、現代の概念は1980年代にまで遡る。特に、スタンフォード大学とカリフォルニア大学バークレー校の2つのプロジェクトが、この概念の普及に最も関係しています。スタンフォード大学のMIPSは成功したMIPSアーキテクチャとして商業化され、バークレー大学のRISCはコンセプト全体に名前を与え、SPARCとして商業化された。この時代のもう一つの成功は、IBMの努力であり、最終的にはIBM POWER命令セットアーキテクチャ、PowerPC、Power ISAにつながった。これらのプロジェクトが成熟するにつれて、1980年代後半から特に1990年代前半にかけては、さまざまな類似の設計が盛んになり、Unixワークステーション市場やレーザープリンタ、ルータなどの組み込みプロセッサに大きな力を発揮するようになりました。

RISC設計には、ARC、Alpha、Am29000、ARM、Atmel AVR、Blackfin、i860、i960、M88000、MIPS、PA-RISC、Power ISA (PowerPCを含む)、RISC-V、SuperH、SPARCなどの多くの種類があります。iPadやAndroidデバイスなどのスマートフォンやタブレットコンピュータでARMアーキテクチャプロセッサが使用されたことで、RISCベースのシステムに幅広いユーザー基盤が提供されました。また、2020年1月現在、TOP500プロジェクトのランキングで世界最速のスーパーコンピュータとなっているSummitなどのスーパーコンピュータにもRISCプロセッサが採用されている[4]。

## インストラクションセットの哲学

命令セット削減型コンピュータ」という言葉によくある誤解は、単に命令を排除しただけで命令セットが小さくなるというものである[2] 。 実際には、RISC の命令セットは年々大きくなり、現在では多くの CISC CPU よりも大きな命令セットを持つものが多い[21] [22] 。 21][22] PowerPCのようなRISCプロセッサの中には、例えばCISCのIBM System/370のような大きな命令セットを持つものもあるが、逆にDECのPDP-8は、命令の多くが複数のメモリアクセスを伴うため、明らかにCISC CPUであることがわかるが、基本命令が8個と拡張命令が数個しかない。

この言葉の「削減」という言葉は、1つの命令を実行するために何十回ものデータメモリサイクルを必要とするCISC CPUの「複雑な命令」と比較して、1つの命令が達成する作業量が削減されているという事実を表現することを意図したものである[23]。 特に、RISCプロセッサは一般的にI/Oとデータ処理のための命令を別々に持っている[引用が必要]。

ロード/ストアアーキテクチャという用語が好まれることがあります。

## 他のアーキテクチャとの比較

一部のCPUは、非常に小さな命令セットを持つように特別に設計されていますが、これらの設計は古典的なRISC設計とは大きく異なるため、ミニマム命令セットコンピュータ(MISC)やトランスポート・トリガ・アーキテクチャ(TTA)などの別の名前が付けられています。

RISCアーキテクチャは、従来、デスクトップPCやコモディティサーバ市場では、x86ベースのプラットフォームが主流のプロセッサアーキテクチャで、ほとんど成功を収めていませんでした。しかし、より高性能なシステム向けにARMベースのプロセッサが開発されているため、これは変わるかもしれない[26]。 Cavium、AMD、QualcommなどのメーカーがARMアーキテクチャをベースにしたサーバプロセッサをリリースしている[27][28]。 ARMは2017年にさらにCrayと提携し、ARMベースのスーパーコンピュータを製造している[29]。 デスクトップでは、MicrosoftはQualcommとの提携の一環として、2017年にQualcomm SnapdragonベースのデバイスでPC版のWindows 10をサポートする予定であることを発表した。これらのデバイスは、x86プロセッサエミュレータを介してx86ベースのWin32ソフトウェアをサポートする予定である[30]。

しかし、デスクトップの分野以外では、ARM RISCアーキテクチャは、スマートフォン、タブレット、および多くの形態の組み込みデバイスで広く使用されている。また、Pentium Pro(P6)以降、Intelがプロセッサに内蔵RISCプロセッサコアを採用しているのも事実である[31]。

初期のRISC設計は現代のCISC設計と大きく異なっていたが、2000年までにはRISCラインの最高性能CPUはCISCラインの最高性能CPUとほとんど区別がつかなくなっていた[32][33][34]。

## AMDもIntelも内部RISC CPUを使っているので、x86をバイパスして内部RISCを直接動かすようなソフトウェアを書くことは可能なのでしょうか？

https://www.quora.com/Since-AMD-and-Intel-both-use-an-internal-RISC-CPU-is-it-possible-to-write-software-that-bypass-x86-and-directly-run-the-internal-RISC

さて、整理してみましょう。現代のx86 CPUは "内部RISC CPU "を持っていません。そのように表現する人は短絡的で、かなり誤解を招くような気がします。

最新の高性能 x86 CPU は、命令を µops (マイクロオペレーション) と呼ばれるものにデコードします。これらのμopsは、プロセッサ内の小さな作業単位を表しています。多くの命令は単一のµopsにデコードされますが、多くのµopsに拡張される命令もあります。場合によっては、1 つの命令が 0µops に拡張されることもあります。一部のプロセッサのマイクロアーキテクチャでは、CPU は可能であれば、特定のμops を融合させることもできます。

元のコードに表示されているレジスタ名も、マシンの実際の物理レジスタを指すように動的に名前が変更されます。x86-64 で表示されている 16 個のレジスタは、はるかに大きな (100 個以上の) ハードウェアレジスタのプールにマップされています。

面白い事実: リネームエンジンは、リネームテーブルを更新して命令を破棄するだけで、いくつかの命令(特にレジスタ間の移動)を0 µopsに拡張できるようにしています。

実際、同じ命令でもレジスタのマッピングが異なる場合があります。これは特にタイトループでは重要で、プロセッサが何度もループを繰り返し、最初の繰り返しが完了する前にデコードしてスケジューリングユニットにμopsを発行してしまうような場合に、先回りして推測することがあります。

入力される命令のμopsへのマッピングは、その命令に遭遇したときのCPUの状態に依存します。マッピングは、実行中のプログラムの正確なフローや、命令が終了する順序に基づいて動的に行われます。

その結果、μopsはプログラミングに適したインターフェースを提供しません。パイプラインのフロントエンドから x86 デコーダとリネームエンジンを取り除き、代わりに生の µops を投入しようとした場合、困難な時間を過ごすことになるでしょう。例えば、レジスタリネームの利点を失うことになるので、まともな性能のループを書くのはかなり難しいでしょう。

* しかし、コンパイラは複数のバージョンのコードを出力したり、実行時に動的に代替案を選択したりすることができます。これは特定のクラスの問題には有効ですが、すべての問題には有効ではありません。また、これはかなり肥大化したオブジェクトコードになります。



# Load–store architecture

[https://en.wikipedia.org/wiki/Load%E2%80%93store_architecture](https://en.wikipedia.org/wiki/Load–store_architecture)

コンピュータ工学では、ロードストアアーキテクチャとは、メモリアクセス（メモリとレジスタ間のロードとストア）とALU演算（レジスタ間でのみ発生する演算）の2つのカテゴリに命令を分割した命令セットアーキテクチャのことである[1]:9-12。

PowerPC, SPARC, RISC-V, ARM, MIPSなどのRISC命令セットアーキテクチャはロードストアアーキテクチャである[1]:9-12.

例えば、ロードストア・アプローチでは、ADD演算のオペランドと宛先の両方がレジスタ内になければなりません。これは、ADD演算のオペランドの一方がメモリ内にあり、他方がレジスタ内にあるようなレジスタ・メモリ・アーキテクチャ(例えば、x86のようなCISC命令セット・アーキテクチャ)とは異なります[1]:9-12

ロードストアアーキテクチャの最も初期の例はCDC 6600であった[1]:54-56 ほとんどすべてのベクトルプロセッサ（多くのGPUを含む[2]）はロードストアアプローチを使用している[3]。

