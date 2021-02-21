#!/bin/bash
####################################
###################################
##
## Performance Tuning toolkit
## by Stefano Trallori#
##
##
##
##
#set -x

typeset -x v_startTimestamp="$(date +%Y%m%d_%H%M%S)"
typeset -x v_hostname="$(hostname)"
typeset -x v_platform="$(uname)"
typeset -x v_platformRelease="$(uname -r)"
typeset -x v_platformBit="$(getconf LONG_BIT)"
typeset -x v_toolkitHome="$( cd "$(dirname "$0")" ; pwd -P )"
typeset -x v_oracleBase="${v_toolkitHome}/oracle"
typeset -x v_tnsAdmin="${v_toolkitHome}/oracle/network/admin"
typeset -x v_contribBase="${v_toolkitHome}/contrib"
typeset -x v_selectedOHPath=
typeset -x v_selectedCHPath=
typeset -x v_OHDir=

f_logMessage () {
  local l_message="${1}"
  local l_severity="${2}"
  local l_timestamp="$(date +%Y%m%d_%H%M%S)"
  local l_returnCode=
  printf "[ %-150s ] [ %s ] [ %s ]\n" "${l_message}" "${l_severity}" "${l_timestamp}"
  l_returnCode=${?}
  return ${l_returnCode}
}

f_getOracleHomes () {
  local l_oracleBase="${v_oracleBase}"
  local l_returnCode=
  local l_index=0
  local l_maxIndex=99
  local l_oracleHomes="$(find "${l_oracleBase}/client/${v_platform}" -type d)"
  local l_oracleHome=
  local l_selectedOH=-1
  declare -a a_oracleHomes=

  while read l_oracleHome; do
    find "${l_oracleHome}" -name *.zip -depth 1 | grep -q . >/dev/null 2>&1
    if [ ${?} -eq 0 ]; then
      a_oracleHomes[${l_index}]="${l_oracleHome}"
      l_maxIndex=${l_index}
      let l_index++
    fi
  done < <(echo "${l_oracleHomes}")
  
  [ ${l_maxIndex} -ne 99 ] && { f_logMessage "Select Oracle Home" "INFO"; } || { f_logMessage "No Oracle Home found in ${v_oracleBase}" "ERROR"; return 1;}
  for (( l_index=0; l_index<=${l_maxIndex}; l_index++ )) 
  do
    echo "[${l_index}] [${a_oracleHomes[${l_index}]}]"
  done
  while true; do
    echo -n "Your choice:   "
    read l_selectedOH
    echo "Your choice was: "${l_selectedOH}
    if [ -z ${l_selectedOH} ] || ! [ ${l_selectedOH} -ge 0  -a ${l_selectedOH} -le ${l_maxIndex} ]; then continue; fi
    eval "v_selectedOHPath=\"${a_oracleHomes[${l_selectedOH}]}\""
    break
  done
  return 0
}

f_installOracleHome () {
  local l_returnCode=
  local l_oracleHome="${1}"
  { cd "${l_oracleHome}"; sh install.sh; cd "${v_toolkitHome}"; } >/dev/null 2>&1
  l_returnCode=${?}
  eval "v_OHDir=\"$(find "${v_selectedOHPath}" -type d -depth 1 | sort | head -1)\""
  return ${l_returnCode}
}

f_setVariables () {
  local l_returnCode=
  eval "PATH=\"${PATH}:${v_OHDir}\""
  eval "TNS_ADMIN=\"${v_tnsAdmin}\""
  return 0
}

f_setAliases () {
  eval "alias vit='vi \"\${TNS_ADMIN}/tnsnames.ora\"'"
  eval "alias visql='vi \"\${TNS_ADMIN}/sqlnet.ora\"'"
}

f_getContribs () {
  local l_contribBase="${v_contribBase}"
  local l_returnCode=
  local l_index=0
  local l_maxIndex=99
  local l_contribHomes="$(find "${l_contribBase}" -type d -depth 2)"
  local l_contribHome=
  local l_selectedCH=-1
  declare -a a_contribHomes=

  while read l_contribHome; do
    a_contribHomes[${l_index}]="${l_contribHome}"
    l_maxIndex=${l_index}
    let l_index++
  done < <(echo "${l_contribHomes}")
   
  a_contribHomes[${l_index}]="quit"
  let l_index++
  let l_maxIndex++

  [ ${l_maxIndex} -ne 99 ] && { f_logMessage "Select contrib home" "INFO"; } || { f_logMessage "No contrib home found in ${v_contribBase}" "ERROR"; return 1;}
  for (( l_index=0; l_index<=${l_maxIndex}; l_index++ ))
  do
    echo "[${l_index}] [${a_contribHomes[${l_index}]}]"
  done
  while true; do
    echo -n "Your choice:   "
    read l_selectedCH
    echo "Your choice was: "${l_selectedCH}
    if [ -z ${l_selectedCH} ] || ! [ ${l_selectedCH} -ge 0  -a ${l_selectedCH} -le ${l_maxIndex} ]; then continue; fi
    if ! [ "${a_contribHomes[${l_selectedCH}]}" == "quit" ]; then
      eval "v_selectedCHPath=\"${a_contribHomes[${l_selectedCH}]}\""
      echo "Installing "${a_contribHomes[${l_selectedCH}]}""
      . "${a_contribHomes[${l_selectedCH}]}"/install/install
    else
      echo "Quitting"
      break
    fi
  done
  return 0
}



f_logMessage "Starting toolkit on ${v_hostname} - ${v_platform} ${v_platformRelease} ${v_platformBit}bit" "INFO"
f_getOracleHomes
f_logMessage "Selected Oracle Home ${v_selectedOHPath}" "INFO"
f_logMessage "Installing Oracle Home ${v_selectedOHPath}" "INFO"
f_installOracleHome "${v_selectedOHPath}"
f_logMessage "Oracle Home installed in ${v_OHDir}" "INFO"
f_logMessage "Setting variables" "INFO"
f_setVariables
f_logMessage "Setting aliases" "INFO"
f_setAliases
f_getContribs
