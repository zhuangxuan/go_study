#!/bin/bash
 
# Tetris Game
# 10.21.2003 xhchen<[email]xhchen@winbond.com.tw[/email]>
 
#APP declaration
APP_NAME="${0##*[\\/]}"
APP_VERSION="1.0"
 
 
#��ɫ����
cRed=1
cGreen=2
cYellow=3
cBlue=4
cFuchsia=5
cCyan=6
cWhite=7
colorTable=($cRed $cGreen $cYellow $cBlue $cFuchsia $cCyan $cWhite)
 
#λ�úʹ�С
iLeft=3
iTop=2
((iTrayLeft = iLeft + 2))
((iTrayTop = iTop + 1))
((iTrayWidth = 10))
((iTrayHeight = 15))
 
#��ɫ����
cBorder=$cGreen
cScore=$cFuchsia
cScoreValue=$cCyan
 
#�����ź�
#����Ϸʹ���������̣�һ�����ڽ������룬һ��������Ϸ���̺���ʾ����;
#��ǰ�߽��յ��������ҵȰ���ʱ��ͨ������߷���signal�ķ�ʽ֪ͨ���ߡ�
sigRotate=25
sigLeft=26
sigRight=27
sigDown=28
sigAllDown=29
sigExit=30
 
#���в�ͬ�ķ���Ķ���
#ͨ����ת��ÿ�ַ������ʾ����ʽ�����м���
box0=(0 0 0 1 1 0 1 1)
box1=(0 2 1 2 2 2 3 2 1 0 1 1 1 2 1 3)
box2=(0 0 0 1 1 1 1 2 0 1 1 0 1 1 2 0)
box3=(0 1 0 2 1 0 1 1 0 0 1 0 1 1 2 1)
box4=(0 1 0 2 1 1 2 1 1 0 1 1 1 2 2 2 0 1 1 1 2 0 2 1 0 0 1 0 1 1 1 2)
box5=(0 1 1 1 2 1 2 2 1 0 1 1 1 2 2 0 0 0 0 1 1 1 2 1 0 2 1 0 1 1 1 2)
box6=(0 1 1 1 1 2 2 1 1 0 1 1 1 2 2 1 0 1 1 0 1 1 2 1 0 1 1 0 1 1 1 2)
#�������з���Ķ��嶼�ŵ�box������
box=(${box0[@]} ${box1[@]} ${box2[@]} ${box3[@]} ${box4[@]} ${box5[@]} ${box6[@]})
#���ַ�����ת����ܵ���ʽ��Ŀ
countBox=(1 2 2 2 4 4 4)
#���ַ�����box�����е�ƫ��
offsetBox=(0 1 3 5 7 11 15)
 
#ÿ���һ���ٶȼ���Ҫ���۵ķ���
iScoreEachLevel=50        #be greater than 7
 
#����ʱ����
sig=0                #���յ���signal
iScore=0        #�ܷ�
iLevel=0        #�ٶȼ�
boxNew=()        #������ķ����λ�ö���
cBoxNew=0        #������ķ������ɫ
iBoxNewType=0        #������ķ��������
iBoxNewRotate=0        #������ķ������ת�Ƕ�
boxCur=()        #��ǰ�����λ�ö���
cBoxCur=0        #��ǰ�������ɫ
iBoxCurType=0        #��ǰ���������
iBoxCurRotate=0        #��ǰ�������ת�Ƕ�
boxCurX=-1        #��ǰ�����x����λ��
boxCurY=-1        #��ǰ�����y����λ��
iMap=()                #��������ͼ��
 
#��ʼ�����б�������Ϊ-1, ��ʾû�з���
for ((i = 0; i < iTrayHeight * iTrayWidth; i++)); do iMap[$i]=-1; done
 
 
#��������Ľ��̵�������
function RunAsKeyReceiver()
{
        local pidDisplayer key aKey sig cESC sTTY
 
        pidDisplayer=$1
        aKey=(0 0 0)
 
        cESC=`echo -ne "\033"`
        cSpace=`echo -ne "\040"`
 
        #�����ն����ԡ���read -s��ȡ�ն˼�ʱ���ն˵����Իᱻ��ʱ�ı䡣
        #�����read -sʱ���򱻲���ɱ�������ܻᵼ���ն˻��ң�
        #��Ҫ�ڳ����˳�ʱ�ָ��ն����ԡ�
        sTTY=`stty -g`
 
        #��׽�˳��ź�
        trap "MyExit;" INT TERM
        trap "MyExitNoSub;" $sigExit
 
        #���ع��
        echo -ne "\033[?25l"
 
 
        while :
        do
                #��ȡ���롣ע-s�����ԣ�-n����һ���ַ���������
                read -s -n 1 key
 
                aKey[0]=${aKey[1]}
                aKey[1]=${aKey[2]}
                aKey[2]=$key
                sig=0
 
                #�ж������˺��ּ�
                if [[ $key == $cESC && ${aKey[1]} == $cESC ]]
                then
                        #ESC��
                        MyExit
                elif [[ ${aKey[0]} == $cESC && ${aKey[1]} == "[" ]]
                then
                        if [[ $key == "A" ]]; then sig=$sigRotate        #<���ϼ�>
                        elif [[ $key == "B" ]]; then sig=$sigDown        #<���¼�>
                        elif [[ $key == "D" ]]; then sig=$sigLeft        #<�����>
                        elif [[ $key == "C" ]]; then sig=$sigRight        #<���Ҽ�>
                        fi
                elif [[ $key == "W" || $key == "w" ]]; then sig=$sigRotate        #W, w
                elif [[ $key == "S" || $key == "s" ]]; then sig=$sigDown        #S, s
                elif [[ $key == "A" || $key == "a" ]]; then sig=$sigLeft        #A, a
                elif [[ $key == "D" || $key == "d" ]]; then sig=$sigRight        #D, d
                elif [[ "[$key]" == "[]" ]]; then sig=$sigAllDown        #�ո��
                elif [[ $key == "Q" || $key == "q" ]]                        #Q, q
                then
                        MyExit
                fi
 
                if [[ $sig != 0 ]]
                then
                        #����һ���̷�����Ϣ
                        kill -$sig $pidDisplayer
                fi
        done
}
 
#�˳�ǰ�Ļָ�
function MyExitNoSub()
{
        local y
 
        #�ָ��ն�����
        stty $sTTY
        ((y = iTop + iTrayHeight + 4))
 
        #��ʾ���
        echo -e "\033[?25h\033[${y};0H"
        exit
}
 
 
function MyExit()
{
        #֪ͨ��ʾ������Ҫ�˳�
        kill -$sigExit $pidDisplayer
 
        MyExitNoSub
}
 
 
#������ʾ����Ϸ���̵�������
function RunAsDisplayer()
{
        local sigThis
        InitDraw
 
        #���ظ����źŵĴ�����
        trap "sig=$sigRotate;" $sigRotate
        trap "sig=$sigLeft;" $sigLeft
        trap "sig=$sigRight;" $sigRight
        trap "sig=$sigDown;" $sigDown
        trap "sig=$sigAllDown;" $sigAllDown
        trap "ShowExit;" $sigExit
 
        while :
        do
                #���ݵ�ǰ���ٶȼ�iLevel��ͬ���趨��Ӧ��ѭ���Ĵ���
                for ((i = 0; i < 21 - iLevel; i++))
                do
                        sleep 0.02
                        sigThis=$sig
                        sig=0
 
                        #����sig�����ж��Ƿ���ܵ���Ӧ���ź�
                        if ((sigThis == sigRotate)); then BoxRotate;        #��ת
                        elif ((sigThis == sigLeft)); then BoxLeft;        #����һ��
                        elif ((sigThis == sigRight)); then BoxRight;        #����һ��
                        elif ((sigThis == sigDown)); then BoxDown;        #����һ��
                        elif ((sigThis == sigAllDown)); then BoxAllDown;        #���䵽��
                        fi
                done
                #kill -$sigDown $$
                BoxDown        #����һ��
        done
}
 
 
#BoxMove(y, x), �����Ƿ���԰��ƶ��еķ����Ƶ�(x, y)��λ��, ����0�����, 1������
function BoxMove()
{
        local j i x y xTest yTest
        yTest=$1
        xTest=$2
        for ((j = 0; j < 8; j += 2))
        do
                ((i = j + 1))
                ((y = ${boxCur[$j]} + yTest))
                ((x = ${boxCur[$i]} + xTest))
                if (( y < 0 || y >= iTrayHeight || x < 0 || x >= iTrayWidth))
                then
                        #ײ��ǽ����
                        return 1
                fi
                if ((${iMap[y * iTrayWidth + x]} != -1 ))
                then
                        #ײ�������Ѿ����ڵķ�����
                        return 1
                fi
        done
        return 0;
}
 
 
#����ǰ�ƶ��еķ���ŵ�����������ȥ,
#�������µķ������ٶȼ���(��һ�η����䵽�ײ�)
function Box2Map()
{
        local j i x y xp yp line
 
        #����ǰ�ƶ��еķ���ŵ�����������ȥ
        for ((j = 0; j < 8; j += 2))
        do
                ((i = j + 1))
                ((y = ${boxCur[$j]} + boxCurY))
                ((x = ${boxCur[$i]} + boxCurX))
                ((i = y * iTrayWidth + x))
                iMap[$i]=$cBoxCur
        done
 
        #��ȥ�ɱ���ȥ����
        line=0
        for ((j = 0; j < iTrayWidth * iTrayHeight; j += iTrayWidth))
        do
                for ((i = j + iTrayWidth - 1; i >= j; i--))
                do
                        if ((${iMap[$i]} == -1)); then break; fi
                done
                if ((i >= j)); then continue; fi
 
                ((line++))
                for ((i = j - 1; i >= 0; i--))
                do
                        ((x = i + iTrayWidth))
                        iMap[$x]=${iMap[$i]}
                done
                for ((i = 0; i < iTrayWidth; i++))
                do
                        iMap[$i]=-1
                done
        done
 
        if ((line == 0)); then return; fi
 
        #������ȥ������line����������ٶȼ�
        ((x = iLeft + iTrayWidth * 2 + 7))
        ((y = iTop + 11))
        ((iScore += line * 2 - 1))
        #��ʾ�µķ���
        echo -ne "\033[1m\033[3${cScoreValue}m\033[${y};${x}H${iScore}         "
        if ((iScore % iScoreEachLevel < line * 2 - 1))
        then
                if ((iLevel < 20))
                then
                        ((iLevel++))
                        ((y = iTop + 14))
                        #��ʾ�µ��ٶȼ�
                        echo -ne "\033[3${cScoreValue}m\033[${y};${x}H${iLevel}        "
                fi
        fi
        echo -ne "\033[0m"
 
 
        #������ʾ��������
        for ((y = 0; y < iTrayHeight; y++))
        do
                ((yp = y + iTrayTop + 1))
                ((xp = iTrayLeft + 1))
                ((i = y * iTrayWidth))
                echo -ne "\033[${yp};${xp}H"
                for ((x = 0; x < iTrayWidth; x++))
                do
                        ((j = i + x))
                        if ((${iMap[$j]} == -1))
                        then
                                echo -ne "  "
                        else
                                echo -ne "\033[1m\033[7m\033[3${iMap[$j]}m\033[4${iMap[$j]}m[]\033[0m"
                        fi
                done
        done
}
 
 
#����һ��
function BoxDown()
{
        local y s
        ((y = boxCurY + 1))        #�µ�y����
        if BoxMove $y $boxCurX        #�����Ƿ��������һ��
        then
                s="`DrawCurBox 0`"        #���ɵķ���Ĩȥ
                ((boxCurY = y))
                s="$s`DrawCurBox 1`"        #��ʾ�µ�����󷽿�
                echo -ne $s
        else
                #�ߵ����, �������������
                Box2Map                #����ǰ�ƶ��еķ�����������������
                RandomBox        #�����µķ���
        fi
}
 
#����һ��
function BoxLeft()
{
        local x s
        ((x = boxCurX - 1))
        if BoxMove $boxCurY $x
        then
                s=`DrawCurBox 0`
                ((boxCurX = x))
                s=$s`DrawCurBox 1`
                echo -ne $s
        fi
}
 
#����һ��
function BoxRight()
{
        local x s
        ((x = boxCurX + 1))
        if BoxMove $boxCurY $x
        then
                s=`DrawCurBox 0`
                ((boxCurX = x))
                s=$s`DrawCurBox 1`
                echo -ne $s
        fi
}
 
 
#���䵽��
function BoxAllDown()
{
        local k j i x y iDown s
        iDown=$iTrayHeight
 
        #����һ����Ҫ���������
        for ((j = 0; j < 8; j += 2))
        do
                ((i = j + 1))
                ((y = ${boxCur[$j]} + boxCurY))
                ((x = ${boxCur[$i]} + boxCurX))
                for ((k = y + 1; k < iTrayHeight; k++))
                do
                        ((i = k * iTrayWidth + x))
                        if (( ${iMap[$i]} != -1)); then break; fi
                done
                ((k -= y + 1))
                if (( $iDown > $k )); then iDown=$k; fi
        done
 
        s=`DrawCurBox 0`        #���ɵķ���Ĩȥ
        ((boxCurY += iDown))
        s=$s`DrawCurBox 1`        #��ʾ�µ������ķ���
        echo -ne $s
        Box2Map                #����ǰ�ƶ��еķ�����������������
        RandomBox        #�����µķ���
}
 
 
#��ת����
function BoxRotate()
{
        local iCount iTestRotate boxTest j i s
        iCount=${countBox[$iBoxCurType]}        #��ǰ�ķ��龭��ת���Բ�������ʽ����Ŀ
 
        #������ת����µ���ʽ
        ((iTestRotate = iBoxCurRotate + 1))
        if ((iTestRotate >= iCount))
        then
                ((iTestRotate = 0))
        fi
 
        #���µ��µ���ʽ, �����ϵ���ʽ(������ʾ)
        for ((j = 0, i = (${offsetBox[$iBoxCurType]} + $iTestRotate) * 8; j < 8; j++, i++))
        do
                boxTest[$j]=${boxCur[$j]}
                boxCur[$j]=${box[$i]}
        done
 
        if BoxMove $boxCurY $boxCurX        #������ת���Ƿ��пռ�ŵ���
        then
                #Ĩȥ�ɵķ���
                for ((j = 0; j < 8; j++))
                do
                        boxCur[$j]=${boxTest[$j]}
                done
                s=`DrawCurBox 0`
 
                #�����µķ���
                for ((j = 0, i = (${offsetBox[$iBoxCurType]} + $iTestRotate) * 8; j < 8; j++, i++))
                do
                        boxCur[$j]=${box[$i]}
                done
                s=$s`DrawCurBox 1`
                echo -ne $s
                iBoxCurRotate=$iTestRotate
        else
                #������ת�����Ǽ���ʹ���ϵ���ʽ
                for ((j = 0; j < 8; j++))
                do
                        boxCur[$j]=${boxTest[$j]}
                done
        fi
}
 
 
#DrawCurBox(bDraw), ���Ƶ�ǰ�ƶ��еķ���, bDrawΪ1, ����, bDrawΪ0, Ĩȥ���顣
function DrawCurBox()
{
        local i j t bDraw sBox s
        bDraw=$1
 
        s=""
        if (( bDraw == 0 ))
        then
                sBox="\040\040"
        else
                sBox="[]"
                s=$s"\033[1m\033[7m\033[3${cBoxCur}m\033[4${cBoxCur}m"
        fi
 
        for ((j = 0; j < 8; j += 2))
        do
                ((i = iTrayTop + 1 + ${boxCur[$j]} + boxCurY))
                ((t = iTrayLeft + 1 + 2 * (boxCurX + ${boxCur[$j + 1]})))
                #\033[y;xH, ��굽(x, y)��
                s=$s"\033[${i};${t}H${sBox}"
        done
        s=$s"\033[0m"
        echo -n $s
}
 
 
#�����µķ���
function RandomBox()
{
        local i j t
 
        #���µ�ǰ�ƶ��ķ���
        iBoxCurType=${iBoxNewType}
        iBoxCurRotate=${iBoxNewRotate}
        cBoxCur=${cBoxNew}
        for ((j = 0; j < ${#boxNew[@]}; j++))
        do
                boxCur[$j]=${boxNew[$j]}
        done
 
 
        #��ʾ��ǰ�ƶ��ķ���
        if (( ${#boxCur[@]} == 8 ))
        then
                #���㵱ǰ����ôӶ�����һ��"ð"����
                for ((j = 0, t = 4; j < 8; j += 2))
                do
                        if ((${boxCur[$j]} < t)); then t=${boxCur[$j]}; fi
                done
                ((boxCurY = -t))
                for ((j = 1, i = -4, t = 20; j < 8; j += 2))
                do
                        if ((${boxCur[$j]} > i)); then i=${boxCur[$j]}; fi
                        if ((${boxCur[$j]} < t)); then t=${boxCur[$j]}; fi
                done
                ((boxCurX = (iTrayWidth - 1 - i - t) / 2))
 
                #��ʾ��ǰ�ƶ��ķ���
                echo -ne `DrawCurBox 1`
 
                #�������һ������û���ţ�Game over!
                if ! BoxMove $boxCurY $boxCurX
                then
                        kill -$sigExit ${PPID}
                        ShowExit
                fi
        fi
 
 
 
        #����ұ�Ԥ��ʾ�ķ���
        for ((j = 0; j < 4; j++))
        do
                ((i = iTop + 1 + j))
                ((t = iLeft + 2 * iTrayWidth + 7))
                echo -ne "\033[${i};${t}H        "
        done
 
        #��������µķ���
        ((iBoxNewType = RANDOM % ${#offsetBox[@]}))
        ((iBoxNewRotate = RANDOM % ${countBox[$iBoxNewType]}))
        for ((j = 0, i = (${offsetBox[$iBoxNewType]} + $iBoxNewRotate) * 8; j < 8; j++, i++))
        do
                boxNew[$j]=${box[$i]};
        done
 
        ((cBoxNew = ${colorTable[RANDOM % ${#colorTable[@]}]}))
 
        #��ʾ�ұ�Ԥ��ʾ�ķ���
        echo -ne "\033[1m\033[7m\033[3${cBoxNew}m\033[4${cBoxNew}m"
        for ((j = 0; j < 8; j += 2))
        do
                ((i = iTop + 1 + ${boxNew[$j]}))
                ((t = iLeft + 2 * iTrayWidth + 7 + 2 * ${boxNew[$j + 1]}))
                echo -ne "\033[${i};${t}H[]"
        done
        echo -ne "\033[0m"
}
 
 
#��ʼ����
function InitDraw()
{
        clear
        RandomBox        #����������飬��ʱ�ұ�Ԥ��ʾ�������з�����
        RandomBox        #������������飬�ұ�Ԥ��ʾ�����еķ��鱻���£�ԭ�ȵķ��齫��ʼ����
        local i t1 t2 t3
 
        #��ʾ�߿�
        echo -ne "\033[1m"
        echo -ne "\033[3${cBorder}m\033[4${cBorder}m"
 
        ((t2 = iLeft + 1))
        ((t3 = iLeft + iTrayWidth * 2 + 3))
        for ((i = 0; i < iTrayHeight; i++))
        do
                ((t1 = i + iTop + 2))
                echo -ne "\033[${t1};${t2}H||"
                echo -ne "\033[${t1};${t3}H||"
        done
 
        ((t2 = iTop + iTrayHeight + 2))
        for ((i = 0; i < iTrayWidth + 2; i++))
        do
                ((t1 = i * 2 + iLeft + 1))
                echo -ne "\033[${iTrayTop};${t1}H=="
                echo -ne "\033[${t2};${t1}H=="
        done
        echo -ne "\033[0m"
 
 
        #��ʾ"Score"��"Level"����
        echo -ne "\033[1m"
        ((t1 = iLeft + iTrayWidth * 2 + 7))
        ((t2 = iTop + 10))
        echo -ne "\033[3${cScore}m\033[${t2};${t1}HScore"
        ((t2 = iTop + 11))
        echo -ne "\033[3${cScoreValue}m\033[${t2};${t1}H${iScore}"
        ((t2 = iTop + 13))
        echo -ne "\033[3${cScore}m\033[${t2};${t1}HLevel"
        ((t2 = iTop + 14))
        echo -ne "\033[3${cScoreValue}m\033[${t2};${t1}H${iLevel}"
        echo -ne "\033[0m"
}
 
 
#�˳�ʱ��ʾGameOVer!
function ShowExit()
{
        local y
        ((y = iTrayHeight + iTrayTop + 3))
        echo -e "\033[${y};0HGameOver!\033[0m"
        exit
}
 
 
#��ʾ�÷�.
function Usage
{
        cat << EOF
Usage: $APP_NAME
Start tetris game.
 
  -h, --help              display this help and exit
      --version           output version information and exit
EOF
}
 
 
#��Ϸ�������������ʼ.
if [[ "$1" == "-h" || "$1" == "--help" ]]; then
        Usage
elif [[ "$1" == "--version" ]]; then
        echo "$APP_NAME $APP_VERSION"
elif [[ "$1" == "--show" ]]; then
        #�����־��в���--showʱ��������ʾ����
        RunAsDisplayer
else
        bash $0 --show&        #�Բ���--show��������������һ��
        RunAsKeyReceiver $!        #����һ�в����Ľ��̵Ľ��̺���Ϊ����
fi