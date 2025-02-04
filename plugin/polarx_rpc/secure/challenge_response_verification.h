/*
 * Copyright (c) 2017, 2018, Oracle and/or its affiliates. All rights reserved.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License, version 2.0,
 * as published by the Free Software Foundation.
 *
 * This program is also distributed with certain software (including
 * but not limited to OpenSSL) that is licensed under separate terms,
 * as designated in a particular file or component or in included license
 * documentation.  The authors of MySQL hereby grant you an additional
 * permission to link the program and your derivative works with the
 * separately licensed software that they have included with MySQL.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License, version 2.0, for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA 02110-1301  USA
 */

#pragma once

#include <string>

#include "crypt_genhash_impl.h"
#include "mysql_com.h"

#include "account_verification_interface.h"
#include "sha256_password_cache_interface.h"

namespace polarx_rpc {

/**
  Class for doing account verification for the challenge response authentication
*/
class Challenge_response_verification : public Account_verification_interface {
 public:
  explicit Challenge_response_verification(
      SHA256_password_cache_interface *cache)
      : k_salt(generate_salt()), m_sha256_password_cache(cache) {}

  const std::string &get_salt() const override { return k_salt; }

 protected:
  const std::string k_salt;
  SHA256_password_cache_interface *m_sha256_password_cache;

  inline std::string generate_salt() {
    std::string salt(SCRAMBLE_LENGTH, '\0');
    ::generate_user_salt(&salt[0], static_cast<int>(salt.size()));
    return salt;
  }
};

}  // namespace polarx_rpc
